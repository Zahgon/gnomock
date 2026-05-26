// Package splunk includes Splunk Enterprise implementation of Gnomock Preset
// interface. This Preset can be passed to gnomock.StartPreset function to
// create a configured Splunk container to use in tests.
//
// Splunk image is relatively heavy (larger than 1.5GB), and its startup time
// is longer than usual. Using this container may make the tests much longer.
package splunk

import (
	"context"
	"net/http"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	// CollectorPort is the name of a port exposed by Splunk Collector.
	CollectorPort string = "collector"

	// APIPort is the name of a port exposed by Splunk API.
	APIPort string = "api"

	// WebPort is the name of a port exposed by Splunk web UI.
	WebPort string = "web"
)

const defaultVersion = "latest"

func init() {
	registry.Register("splunk", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gnomock Splunk preset. This preset includes a Splunk
// specific healthcheck function, default Splunk image and ports, and allows to
// optionally ingest initial logs.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of Splunk.
type P struct {
	Values        []Event `json:"values"`
	ValuesFile    string  `json:"values_file"`
	AcceptLicense bool    `json:"accept_license"`
	AdminPassword string  `json:"admin_password"`
	Version       string  `json:"version"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func healthcheck(password string) gnomock.HealthcheckFunc {
	_ = "STUB: not implemented"
	return *new(gnomock.HealthcheckFunc)
}

func checkAPI(ctx context.Context, c *gnomock.Container, password string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkHEC(ctx context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func insecureClient() http.Client { _ = "STUB: not implemented"; return *new(http.Client) }

//nolint:gosec
