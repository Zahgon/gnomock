package splunk

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithVersion sets splunk version (see
// https://hub.docker.com/r/splunk/splunk/tags) for available versions
func WithVersion(version string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValues initializes Splunk with the provided values as log entries.
func WithValues(vs []Event) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValuesFile sets file name to use as a source of initial events to
// ingest. These events are ingested first, followed by any other events sent
// using WithValues.
func WithValuesFile(file string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLicense lets the user choose to accept Splunk enterprise license
// (see more at https://hub.docker.com/_/splunk-enterprise). Failure to accept
// the license will prevent Splunk container from starting.
func WithLicense(accept bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPassword sets admin password in Splunk container. Use this password to
// connect to the container when it is ready. Note that Splunk has password
// requirements. Failure to meet those will prevent the container from starting
// (see defaults at
// https://docs.splunk.com/Documentation/Splunk/latest/Security/Configurepasswordsinspecfile)
func WithPassword(pass string) Option { _ = "STUB: not implemented"; return *new(Option) }
