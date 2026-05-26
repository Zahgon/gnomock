package vault

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithVersion sets image version.
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthToken sets authentication (root) token to be used to connect to this
// container.
func WithAuthToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuth enables new vault authorizations endpoints.
func WithAuth(auth []Auth) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPolicies configures vault with the provided policies.
func WithPolicies(policies []Policy) Option { _ = "STUB: not implemented"; return *new(Option) }
