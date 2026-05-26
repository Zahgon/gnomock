// Package influxdb includes InfluxDB implementation of Gnomock Preset
// interface. This Preset can be passed to gnomock.Start() function to create a
// configured InfluxDB container to use in tests.
//
// Currently this preset only supports InfluxDB 2.x. If you are interested in
// earlier versions, please consider opening an issue.
package influxdb

import (
	"context"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultVersion   = "2.7.6-alpine"
	defaultPort      = 8086
	defaultUsername  = "gnomock"
	defaultPassword  = "gnomock-password"
	defaultOrg       = "gnomock-org"
	defaultBucket    = "gnomock-bucket"
	defaultAuthToken = "gnomock-influxdb-token" //nolint:gosec
)

func init() {
	registry.Register("influxdb", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock InfluxDB preset. This preset includes a
// InfluxDB specific healthcheck function and default InfluxDB image and port.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation for InfluxDB.
type P struct {
	Version   string `json:"version"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Org       string `json:"org"`
	Bucket    string `json:"bucket"`
	AuthToken string `json:"auth_token"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) healthcheck(ctx context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}
