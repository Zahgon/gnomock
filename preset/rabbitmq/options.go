package rabbitmq

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithUser creates a new superuser with the provided credentials in the
// container.
func WithUser(user, password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVersion sets image version. See
// https://hub.docker.com/_/rabbitmq/?tab=tags for a list of available tags.
// Use ones with "management" to enable RabbitMQ management plugin, and get the
// mapped port using `container.Port(rabbitmq.ManagementPort)`.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessages makes sure that these messages can be consumed during the test
// once the container is ready.
func WithMessages(messages ...Message) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMessagesFile allows to load messages to be sent into RabbitMQ from one
// or multiple files.
func WithMessagesFile(file string) Option { _ = "STUB: not implemented"; return *new(Option) }
