package kafka

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTopics makes sure that the provided topics are available when Kafka is
// up and running.
// Both topics from WithTopics and WithTopicConfigs will be added to Kafka.
func WithTopics(topics ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTopicConfigs makes sure that the provided topics with the given configs are available when Kafka is
// up and running. Unlike WithTopics, this allows to also set partition count.
// Both topics from WithTopics and WithTopicConfigs will be added to Kafka.
func WithTopicConfigs(topics ...TopicConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessages makes sure that these messages can be consumed during the test
// once the container is ready.
func WithMessages(messages ...Message) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessagesFile allows to load messages to be sent into Kafka from one or
// multiple files.
func WithMessagesFile(files string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSchemaRegistry makes the container wait for the schema registry port to
// become available. Note that it takes longer to setup schema registry than
// the broker itself. Gnomock will not wait for the registry by default, but it
// may become available eventually.
func WithSchemaRegistry() Option { _ = "STUB: not implemented"; return *new(Option) }
