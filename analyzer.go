package terraformanalyzer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"

	"gopkg.in/src-d/lookout-sdk.v0/pb"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hclwrite"
	"gopkg.in/src-d/go-log.v1"
)

// this regex checks if this is a terraform hcl file name
var terraformFileRegex = regexp.MustCompile(`\.tf$`)

type Analyzer struct {
	DataClient pb.DataClient
	Version    string
}

func (a Analyzer) NotifyReviewEvent(ctx context.Context, review *pb.ReviewEvent) (*pb.EventResponse, error) {
	log.Infof("got review request %v", review)

	changes, err := a.DataClient.GetChanges(ctx, &pb.ChangesRequest{
		Head:            &review.Head,
		Base:            &review.Base,
		WantContents:    true,
		WantLanguage:    false,
		WantUAST:        false,
		ExcludeVendored: true,
		IncludePattern:  `\.tf$`,
	})

	if err != nil {
		log.Errorf(err, "GetChanges from DataServer failed")
		return nil, err
	}

	var comments []*pb.Comment
	hadFiles := map[string]bool{}

	for {
		change, err := changes.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Errorf(err, "GetChanges from DataServer failed")
			return nil, err
		}

		log.Infof("analyzing '%s'", change.Head.Path)

		if change.Head == nil {
			log.Infof("ignoring deleted '%s'", change.Base.Path)
			continue
		}

		if _, hasAnalyzed := hadFiles[change.Head.Path]; hasAnalyzed {
			log.Infof("ignoring already analyzed '%s'", change.Head.Path)
			continue
		}

		hadFiles[change.Head.Path] = true

		// run the file through the HCL syntax parser
		_, syntaxDiags := hclsyntax.ParseConfig(change.Head.Content, change.Head.Path, hcl.Pos{Line: 1, Column: 1})
		if syntaxDiags.HasErrors() {
			comments = append(comments, &pb.Comment{
				File: change.Head.Path,
				Line: 0,
				Text: fmt.Sprintf("HCL errored on fomatting:\n%s", syntaxDiags),
			})
			continue
		}

		formatted := hclwrite.Format(change.Head.Content)
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
