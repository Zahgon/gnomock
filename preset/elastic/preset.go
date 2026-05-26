// Package elastic provides a Gnomock Preset for Elasticsearch.
package elastic

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultVersion = "8.13.0"
	defaultPort    = 9200
)

func init() {
	registry.Register("elastic", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gnomock Elasticsearch preset. This preset includes an
// Elasticsearch specific healthcheck function and default Elasticsearch image
// and port.
//
// By default, version 8.7.0 is used.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of Elasticsearch.
type P struct {
	Version string   `json:"version"`
	Inputs  []string `json:"input_files"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) healthcheck(_ context.Context, c *gnomock.Container) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) initf(ctx context.Context, c *gnomock.Container) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) ingestSeedFiles(ctx context.Context, client *elasticsearch.Client) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *P) totalDocCount(client *elasticsearch.Client) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *P) ingestFile(fName string, client *elasticsearch.Client) (docCount int, err error) {
	_ = "STUB: not implemented"
	return 0,
		// nolint:gosec
		nil
}

func (p *P) ingestData(index string, bs []byte, client *elasticsearch.Client) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }
