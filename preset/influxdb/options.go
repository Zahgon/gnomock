package influxdb

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUser sets super-user name and password for this container.
func WithUser(username, password string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOrg sets organization name for this container.
func WithOrg(org string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBucket sets the initial bucket name for this container.
func WithBucket(bucket string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthToken sets authentication token to be used to connect to this
// container.
func WithAuthToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }
