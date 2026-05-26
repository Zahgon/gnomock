// Package cockroachdb includes CockroachDB implementation of Gnomock Preset interface.
// This Preset can be passed to gnomock.Start() function to create a configured
// CockroachDB container to use in tests.
//
// Containers created with this preset use `root` user without a password for
// authentication. There is currently no way to setup an initial user at create
// time.
//
// By default, a new database "mydb" is created, and all the provided queries
// are executed against it.
package cockroachdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultVersion  = "v25.1.6"
	defaultPort     = 26257
	defaultDatabase = "mydb"
)

func init() {
	registry.Register("cockroachdb", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock CockroachDB preset. This preset includes a CockroachDB
// specific healthcheck function and default CockroachDB image and port.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation for CockroachDB.
type P struct {
	Version      string   `json:"version"`
	DB           string   `json:"db"`
	Queries      []string `json:"queries"`
	QueriesFiles []string `json:"queries_files"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func healthcheck(_ context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }

// nolint:gosec

func connect(c *gnomock.Container, db string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
