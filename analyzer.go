package terraformanalyzer

import (
	"bytes"
	"context"
	"fmt"
	"regexp"

	"github.com/src-d/lookout"

	"github.com/hashicorp/hcl/hcl/printer"
	"gopkg.in/src-d/go-log.v1"
	"gopkg.in/src-d/lookout-sdk.v0/pb"
)

// this regex checks if this is a terraform hcl file name
var terraformFileRegex = regexp.MustCompile(`\.tf$`)

type Analyzer struct {
	DataClient *lookout.DataClient
	Version    string
}

func (a Analyzer) NotifyReviewEvent(ctx context.Context, review *pb.ReviewEvent) (*pb.EventResponse, error) {
	log.Infof("got review request %v", review)

	changes, err := a.DataClient.GetChanges(ctx, &pb.ChangesRequest{
		Head:            &review.Head,
		Base:            &review.Base,
		WantContents:    true,
		WantUAST:        false,
		ExcludeVendored: true,
	})
	if err != nil {
		log.Errorf(
			err, "GetChanges from DataServer failed")
	}

	var comments []*pb.Comment
	hadFiles := map[string]bool{}
	for changes.Next() {
		change := changes.Change()

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

	return &pb.EventResponse{AnalyzerVersion: a.Version, Comments: comments}, nil
}

func (a Analyzer) NotifyPushEvent(context.Context, *pb.PushEvent) (*pb.EventResponse, error) {
	return &pb.EventResponse{}, nil
}
