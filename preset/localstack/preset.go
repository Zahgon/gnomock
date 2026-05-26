// Package localstack provides a Gnomock Preset for localstack project
// (https://github.com/localstack/localstack). It allows to easily setup local
// AWS stack for testing
package localstack

import (
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	webPort = "web"

	// APIPort should be used to configure AWS SDK endpoint.
	APIPort = "api"
)

const defaultVersion = "0.14.0"

func init() {
	registry.Register("localstack", func() gnomock.Preset { return &P{} })
}

// Preset creates a new localstack preset to use with gnomock.Start. See
// package docs for a list of exposed ports and services. It is legal to not
// provide any services using WithServices options, but in such case a new
// localstack container will be useless.
//
// This Preset cannot be used with localstack image prior to 0.11.0.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset localstack implementation.
type P struct {
	Services []Service `json:"services"`
	S3Path   string    `json:"s3_path"`
	Version  string    `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) healthcheck(services []string) gnomock.HealthcheckFunc {
	_ = "STUB: not implemented"
	return *new(gnomock.HealthcheckFunc)
}

// available status was added in 0.13.0: it allows to lazy load the
// services after the first request

// healthCheckAddress returns the address of `/health` endpoint of a running
// localstack container. Before version 0.11.3, the endpoint was available at
// port 8080. In 0.11.3, the endpoint was moved to the default port (4566). In
// 1.3.0, the endpoint was moved under `_localstack` prefix.
func (p *P) healthCheckAddress(c *gnomock.Container) string { _ = "STUB: not implemented"; return "" }

type healthResponse struct {
	Services map[string]string `json:"services"`
}

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }
