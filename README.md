Terraform (fmt) Lookout Analyzer
================================

This is a [lookout](https://github.com/src-d/lookout) analyzer that checks if your PR has been Terraform fmt'ed when submitting it.

## Testing
To test this analyzer:
1) Start the analyzer server `go run *.go`
2) Get `lookout-sdk` (https://github.com/src-d/lookout/releases)
3) In a test repository run `./lookout-sdk review ipv4://localhost:2020`
