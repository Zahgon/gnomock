// Package azurite provides a Gnomock Preset for azurite project
// It allows to easily setup local
// Blobstorage for testing
package azurite

import (
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	BlobServicePort  = "blob"
	QueueServicePort = "queue"
	TableServicePort = "table"
)

const (
	ConnectionStringFormat = "DefaultEndpointsProtocol=http;AccountName=%s;AccountKey=%s;BlobEndpoint=http://%s/%s;"
	AccountName            = "devstoreaccount1"
	AccountKey             = "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw=="
	DefaultVersion         = "3.34.0"
)

func init() {
	registry.Register("azurite", func() gnomock.Preset { return &P{} })
}

// Preset creates a new azurite preset to use with gnomock.Start. See
// package docs for a list of exposed ports.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset blobstorage implementation.
type P struct {
	BlobstorePath string `json:"blobstore_path"`
	Version       string `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) healthcheck() gnomock.HealthcheckFunc {
	_ = "STUB: not implemented"
	return *new(gnomock.HealthcheckFunc)
}

// needs implementation. unfortunately azurite does not offer a health endpoint
// a netcat p.healtCheckAddress(c) -z could help alternatively

// healthCheckAddress returns the address of azurite blobstorage endpoint of a running
// azurite container.
func (p *P) healthCheckAddress(c *gnomock.Container) string { _ = "STUB: not implemented"; return "" }

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }
