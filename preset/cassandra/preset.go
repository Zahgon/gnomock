// Package cassandra includes Cassandra implementation of Gnomock Preset
// interface. This Preset can be passed to gnomock.Start() function to create a
// configured Cassandra container to use in tests.
//
// Cassandra containers always use cassandra/cassandra username/password pair,
// it is currently not possible to use different values.
package cassandra

import (
	"context"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

// By default, Cassandra containers will use these values.
const (
	DefaultUser     = "cassandra"
	DefaultPassword = "cassandra"

	defaultVersion = "3"
	defaultPort    = 9042
)

func init() {
	registry.Register("cassandra", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock Cassandra preset. This preset includes a
// Cassandra specific healthcheck function and default Cassandra image and
// port.
//
// Containers created using this preset should be accessed using
// cassandra/cassandra username/password pair.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation for Cassandra.
type P struct {
	Version string `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) healthcheck(_ context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}
