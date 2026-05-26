package k3s

// Option is an optional configuration of this Gnomock preset. Use available
// Options to configure the container.
type Option func(*P)

// WithVersion sets image version.
func WithVersion(version string) Option {
	_ = "STUB: not implemented"

	// support original orlangure/k3s versions with rancher/k3s by
	// appending '-k3s1' to them.
	return *new(Option)
}

// --disable flag only exists in k3s versions above v1.17.0.

// Deprecated: WithPort allows to use a custom port for k3s API access instead
// of the default one. If no custom port is provided, port 48443 is used
// instead.
//
// Please make sure that whichever port you choose to use (including the
// default) is available on the host system. Otherwise this container won't
// start.
//
// This option and its affects has been kept as is for backward compatibility.
// We recommend using `WithDynamicPort()` instead as it does not require a
// static port to be available on the host.
func WithPort(port int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDynamicPort configures the preset to find and use a dynamic (free) port
// for k3s API access from the host. The kubeconfig is replaced with the
// container host port so local Kubernetes clients will still work.
//
// This is preferred to `WithPort` as it does not require a specific host port
// to be available.
func WithDynamicPort() Option { _ = "STUB: not implemented"; return *new(Option) }
