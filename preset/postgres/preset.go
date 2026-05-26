// Package postgres provides a Gnomock Preset for PostgreSQL database.
package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultUser     = "postgres"
	defaultPassword = "password"
	defaultDatabase = "postgres"
	defaultSSLMode  = "disable"
	defaultPort     = 5432
	defaultVersion  = "16.2"
)

func init() {
	registry.Register("postgres", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock Postgres preset. This preset includes a Postgres
// specific healthcheck function, default Postgres image and port, and allows to
// optionally set up initial state.
//
// By default, this preset uses `postgres` user with `password` password, with
// default database `postgres`. Default PostgresQL version is 12.5.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of PostgreSQL database.
type P struct {
	DB           string   `json:"db"`
	Queries      []string `json:"queries"`
	QueriesFiles []string `json:"queries_files"`
	User         string   `json:"user"`
	Password     string   `json:"password"`
	Timezone     string   `json:"timezone"`
	Version      string   `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) healthcheck(_ context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) executeQueries(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

// nolint:gosec

func connect(c *gnomock.Container, db string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
