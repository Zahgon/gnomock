package memcached

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithValues initializes Memcached with the provided key/value pairs. These values
// never expire. Only strings are supported.
func WithValues(vs map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithByteValues initializes Memcached with the provided key/value paris. These values
// never expire. Only byte slices are supported.
func WithByteValues(vs map[string][]byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }
