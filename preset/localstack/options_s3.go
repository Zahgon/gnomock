package localstack

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/orlangure/gnomock"
)

// WithS3Files sets up S3 service running in localstack with the contents of
// `path` directory. The first level children of `path` must be directories,
// their names will be used to create buckets. Below them, all the files in any
// other directories, these files will be uploaded as-is.
//
// For example, if you put your test files in testdata/my-bucket/dir/, Gnomock
// will create "my-bucket" for you, and pull "dir" with all its contents into
// this bucket.
//
// This function does nothing if you don't provide localstack.S3 as one of the
// services in WithServices.
func WithS3Files(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (p *P) initS3(c *gnomock.Container) error { _ = "STUB: not implemented"; return nil }

func (p *P) createBuckets(svc *s3.Client) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create buckets from top-level folders under `path`

func (p *P) createBucket(svc *s3.Client, bucket string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) uploadFiles(svc *s3.Client, buckets []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) uploadFile(svc *s3.Client, bucket, file string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec
