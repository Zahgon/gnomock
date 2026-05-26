// Package mssql provides a Gnomock Preset for Microsoft SQL Server database
package mssql

import (
	"context"
	"database/sql"

	_ "github.com/microsoft/go-mssqldb" // mssql driver
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	masterDB        = "master"
	defaultPassword = "Gn0m!ck~"
	defaultDatabase = "mydb"
	defaultPort     = 1433
	defaultVersion  = "2019-latest"
)

func init() {
	registry.Register("mssql", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock Microsoft SQL Server preset. This preset
// includes a mssql specific healthcheck function, default mssql image and
// port, and allows to optionally set up initial state.
//
// When used without any configuration, it uses `mydb` database, and `Gn0m!ck~`
// administrator password (user: `sa`). You must accept EULA to use this image
// (`WithLicense` option). By default, version `2019-latest` is used.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of Microsoft SQL Server database.
type P struct {
	DB           string   `json:"db"`
	Password     string   `json:"password"`
	Queries      []string `json:"queries"`
	QueriesFiles []string `json:"queries_files"`
	License      bool     `json:"license"`
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

// nolint:gosec

func (p *P) connect(addr, db string) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }
