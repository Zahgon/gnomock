// Package mongo includes mongo implementation of Gnomock Preset interface.
// This Preset can be passed to gnomock.StartPreset function to create a
// configured mongo container to use in tests
package mongo

import (
	"context"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
	mongodb "go.mongodb.org/mongo-driver/mongo"
)

const defaultVersion = "5.0"

func init() {
	registry.Register("mongo", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock MongoDB preset. This preset includes a MongoDB
// specific healthcheck function, default MongoDB image and port, and allows to
// optionally set up initial state.
//
// By default, this preset uses MongoDB 4.4.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of MongoDB.
type P struct {
	DataPath string `json:"data_path"`
	User     string `json:"user"`
	Password string `json:"password"`
	Version  string `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) initf(ctx context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) useCustomUser() bool { _ = "STUB: not implemented"; return false }

func (p *P) setupDB(client *mongodb.Client, dirName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) setupCollection(client *mongodb.Client, dirName, dataFileName string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func healthcheck(ctx context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}
