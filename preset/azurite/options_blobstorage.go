package azurite

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/orlangure/gnomock"
)

// WithBlobstorageFiles sets up Blobstorage service running in azurite with the contents of
// `path` directory. The first level children of `path` must be directories,
// their names will be used to create containers. Below them, all the files in any
// other directories, these files will be uploaded as-is.
//
// For example, if you put your test files in testdata/my-container/dir/, Gnomock
// will create "my-container" for you, and pull "dir" with all its contents into
// this container.
func WithBlobstorageFiles(path string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (p *P) initBlobstorage(c *gnomock.Container) error { _ = "STUB: not implemented"; return nil }

func (p *P) createContainer(ctx context.Context, azblobClient *azblob.Client) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create containers from top-level folders under `path`

func (p *P) createContainers(ctx context.Context, azblobClient *azblob.Client, containerName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) uploadFiles(ctx context.Context, azblobClient *azblob.Client, containerNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) uploadFile(ctx context.Context, azblobClient *azblob.Client, containerName, file string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec
