package redis

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithValues initializes Redis with the provided key/value pairs. These values
// never expire. See go-redis/redis package for information on supported value
// types.
func WithValues(vs map[string]interface{}) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }
