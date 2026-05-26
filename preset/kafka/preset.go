// Package kafka provides a Gnomock Preset for Kafka.
//
// This preset cannot be used in parallel tests due to Kafka's port binding
// limitations. See https://github.com/orlangure/gnomock/issues/1038 for more
// details.
package kafka

import (
	"context"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
	"github.com/segmentio/kafka-go"
)

// The following ports are exposed by this preset:.
const (
	BrokerPort         = "broker"
	ZooKeeperPort      = "zookeeper"
	WebPort            = "web"
	SchemaRegistryPort = "registry"
)

const (
	defaultVersion     = "3.6.1-L0"
	brokerPort         = 49092
	zookeeperPort      = 2181
	webPort            = 3030
	schemaRegistryPort = 8081
)

// Message is a single message sent to Kafka.
type Message struct {
	Topic string `json:"topic"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Time  int64  `json:"time"`
}

func init() {
	registry.Register("kafka", func() gnomock.Preset { return &P{} })
}

// Preset creates a new Gmomock Kafka preset. This preset includes a
// Kafka specific healthcheck function and default Kafka image and ports.
//
// Kafka preset uses a constant broker port number (49092) instead of
// allocating a random unoccupied port on every run. Please make sure this port
// is available when using this preset.
//
// By default, this preset uses `lensesio/fast-data-dev` docker image with
// version `2.5.1-L0` (version can be changed using `WithVersion`).
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

type TopicConfig struct {
	Topic         string
	NumPartitions int
}

// P is a Gnomock Preset implementation of Kafka.
type P struct {
	Version           string    `json:"version"`
	Topics            []string  `json:"topics"`
	Messages          []Message `json:"messages"`
	MessagesFiles     []string  `json:"messages_files"`
	UseSchemaRegistry bool      `json:"use_schema_registry"`

	TopicConfigs []TopicConfig `json:"topic_configs"`
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) healthcheck(ctx context.Context, c *gnomock.Container) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) healthcheckRegistry(ctx context.Context, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

func (p *P) initf(ctx context.Context, c *gnomock.Container) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *P) ingestMessageFiles(ctx context.Context, c *gnomock.Container, conn *kafka.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// cannot set more; cluster has just 1 node

// nolint:gosec
func (p *P) loadMessagesFromFile(fName string) (msgs []Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *P) connect(c *gnomock.Container) (*kafka.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nolint: lll
func (p *P) sendMessagesIntoTopic(ctx context.Context, brokerAddr, topic string, messages []Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}
