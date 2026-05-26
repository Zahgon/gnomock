// Package mariadb provides a Gnomock Preset for MariaDB database
package mariadb

import (
	"context"
	"database/sql"
	"sync"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultUser     = "gnomock"
	defaultPassword = "gnoria"
	defaultDatabase = "mydb"
	defaultPort     = 3306
	defaultVersion  = "10.5.8"
)

var setLoggerOnce sync.Once

func init() {
	registry.Register("mariadb", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock MariaDB preset. This preset includes a MariaDB
// specific healthcheck function, default MariaDB image and port, and allows to
// optionally set up initial state.
//
// When used without any configuration, it creates a superuser `gnomock` with
// password `gnoria`, and `mydb` database. Default MariaDB version is 10.5.8.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of MariaDB database.
type P struct {
	DB           string   `json:"db"`
	User         string   `json:"user"`
	Password     string   `json:"password"`
	Queries      []string `json:"queries"`
	QueriesFiles []string `json:"queries_files"`
	Version      string   `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

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
