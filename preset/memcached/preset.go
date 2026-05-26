// Package memcached includes Memcached implementation of Gnomock Preset interface.
// This Preset can be passed to gnomock.StartPreset function to create a
// configured Memcached container to use in tests
package memcached

import (
	"context"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const defaultVersion = "1.6.9"

func init() {
	registry.Register("memcached", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock Memcached preset. This preset includes a Memcached
// specific healthcheck function, default Memcached image and port, and allows to
// optionally set up initial state.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation for Memcached storage.
type P struct {
	Values     map[string]string `json:"values"`
	ByteValues map[string][]byte `json:"byteValues"`
	Version    string            `json:"version"`
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
