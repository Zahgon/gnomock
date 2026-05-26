// Package rabbitmq provides a Gnomock Preset for RabbitMQ.
package rabbitmq

import (
	"context"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ManagementPort is a name of the port exposed by RabbitMQ management plugin.
// This port is only available when an appropriate version of RabbitMQ docker
// image is used. See `Preset` docs for more info.
const ManagementPort = "management"

const (
	defaultUser     = "guest"
	defaultPassword = "guest"
	defaultVersion  = "3.8.9"
	defaultPort     = 5672
	managementPort  = 15672
)

// Message is a single message sent to RabbitMQ.
type Message struct {
	Queue       string `json:"queue"`
	ContentType string `json:"content_type"`
	StringBody  string `json:"string_body"`
	Body        []byte `json:"body"`
}

func init() {
	registry.Register("rabbitmq", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock RabbitMQ preset. This preset includes a
// RabbitMQ specific healthcheck function and default RabbitMQ image and port.
//
// By default, this preset does not use RabbitMQ Management plugin. To enable
// it, use one of the management tags with `WithVersion` option. Management
// port will be accessible using `container.Port(rabbitmq.ManagementPort)`. See
// https://hub.docker.com/_/rabbitmq/?tab=tags for a list of available tags.
//
// When used without specifying username/password, default ones are used:
// guest/guest. Default version for this preset is 3.8.9.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of RabbitMQ.
type P struct {
	User          string    `json:"user"`
	Password      string    `json:"password"`
	Version       string    `json:"version"`
	Messages      []Message `json:"messages"`
	MessagesFiles []string  `json:"messages_files"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) healthcheck(ctx context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

// any non-err response is valid, it is most likely 401 Unauthorized

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) initf(ctx context.Context, c *gnomock.Container) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) ingestMessages(ctx context.Context, conn *amqp.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) loadFiles() error { _ = "STUB: not implemented"; return nil }

func declareQueues(ch *amqp.Channel, qs []string) error { _ = "STUB: not implemented"; return nil }

func (p *P) isManagement() bool { _ = "STUB: not implemented"; return false }

// nolint:gosec
func (p *P) loadMessagesFromFile(fName string) (msgs []Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *P) connect(c *gnomock.Container) (*amqp.Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *P) sendMessagesIntoQueue(ctx context.Context, ch *amqp.Channel, q string, msgs []Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}
