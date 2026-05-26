package azurite

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }
