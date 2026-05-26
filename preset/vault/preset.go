// Package vault includes vault implementation of Gnomock Preset
// interface. This Preset can be passed to gnomock.Start() function to create a
// configured vault container to use in tests.
package vault

import (
	"context"

	"github.com/hashicorp/vault/api"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
)

const (
	defaultVersion   = "1.13.0"
	defaultPort      = 8200
	defaultAuthToken = "gnomock-vault-token" //nolint:gosec
)

func init() {
	registry.Register("vault", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock Vault preset. This preset includes a
// vault specific healthcheck function and default vault image and port.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation for vault.
type P struct {
	Version   string   `json:"version"`
	AuthToken string   `json:"auth_token"`
	Auth      []Auth   `json:"auth"`
	Policies  []Policy `json:"policies"`
}

// Auth represents a vault authorization.
type Auth struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

// Policy is a vault policy.
type Policy struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }

func (p *P) healthcheck(_ context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Client creates a configured vault client for the provided container and token.
func Client(c *gnomock.Container, token string) (*api.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateToken creates an additional access token with the provided policies. Use the same password you provided
// with the WithAuthToken option.
func CreateToken(c *gnomock.Container, rootToken string, policies ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
