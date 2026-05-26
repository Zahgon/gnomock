// Package testutil includes utilities used in test code of other packages.
package testutil

import (
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

// ListContainerByID returns a list of Containers for the given container id.
func ListContainerByID(cli *client.Client, id string) ([]container.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
