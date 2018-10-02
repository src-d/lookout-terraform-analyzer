package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"

	"github.com/hashicorp/hcl/hcl/printer"
	"google.golang.org/grpc"
	"gopkg.in/src-d/go-log.v1"
	"gopkg.in/src-d/lookout-sdk.v0/pb"
)

// this regex checks if this is a terrafrom hcl file name
var terraformFileRegex = regexp.MustCompile(`\.tf$`)

type analyzer struct{}

func (*analyzer) NotifyReviewEvent(ctx context.Context, review *pb.ReviewEvent) (*pb.EventResponse, error) {
	log.Infof("got review request %v", review)

	conn, err := grpc.Dial(dataSrvAddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf(err, "failed to connect to DataServer at %s", dataSrvAddr)
	}
	defer conn.Close()

	dataClient := pb.NewDataClient(conn)
	changes, err := dataClient.GetChanges(ctx, &pb.ChangesRequest{
		Head:            &review.Head,
		Base:            &review.Base,
		WantContents:    true,
		WantUAST:        false,
		ExcludeVendored: true,
	})
	if err != nil {
		log.Errorf(
			err, "GetChanges from DataServer %s failed", dataSrvAddr)
	}

	var comments []*pb.Comment
	hadFiles := map[string]bool{}
	for {
		change, err := changes.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf(err, "GetChanges from DataServer %s failed", dataSrvAddr)
		}

		log.Infof("analyzing '%s'", change.Head.Path)

		if !terraformFileRegex.MatchString(change.Head.Path) {
			log.Infof("'%s' is not an HCL file", change.Head.Path)
			continue
		}

		if _, hasAnalyzed := hadFiles[change.Head.Path]; hasAnalyzed {
			log.Infof("Already analyzed '%s'", change.Head.Path)
			continue
		}

		hadFiles[change.Head.Path] = true

		// run the file through the HCL formatter
		formatted, err := printer.Format(change.Head.Content)
		if err != nil {
			log.Infof("HCL errored on fomatting '%s' with error: %s", change.Head.Path, err)
			continue
		}

		// check if changes have been made
		if !bytes.Equal(change.Head.Content, formatted) {
			comments = append(comments, &pb.Comment{
				File: change.Head.Path,
				Line: 0,
				Text: fmt.Sprintf("This file is not Terraform fmt'ed"),
			})
		}
	}

	return &pb.EventResponse{AnalyzerVersion: version, Comments: comments}, nil
}

func (*analyzer) NotifyPushEvent(context.Context, *pb.PushEvent) (*pb.EventResponse, error) {
	return &pb.EventResponse{}, nil
}
