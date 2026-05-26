// Package mysql provides a Gnomock Preset for MySQL database.
//
// This preset uses different docker images for amd64 and arm64 architectures.
// Even though the versions should be compatible, you should make sure that the
// tag being used exists in the following repositories:
//
// amd64: https://hub.docker.com/_/mysql/
// arm64: https://hub.docker.com/r/mysql/mysql-server
package mysql

import (
	"context"
	"database/sql"
	"sync"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultUser     = "gnomock"
	defaultPassword = "gnomick"
	defaultDatabase = "mydb"
	defaultPort     = 3306
	defaultVersion  = "8.0.22"
)

var setLoggerOnce sync.Once

func init() {
	registry.Register("mysql", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock MySQL preset. This preset includes a MySQL
// specific healthcheck function, default MySQL image and port, and allows to
// optionally set up initial state.
//
// When used without any configuration, it creates a superuser `gnomock` with
// password `gnomick`, and `mydb` database. Default MySQL version is 8.0.22.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of MySQL database.
type P struct {
	DB           string   `json:"db"`
	User         string   `json:"user"`
	Password     string   `json:"password"`
	Queries      []string `json:"queries"`
	QueriesFiles []string `json:"queries_files"`
	Version      string   `json:"version"`
}

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

// err is always nil for non-nil logger

func (p *P) healthcheck(_ context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }

// nolint:gosec

func (p *P) connect(addr string) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }
