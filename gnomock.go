// Package gnomock contains a framework to set up temporary docker containers
// for integration and end-to-end testing of other applications. It handles
// pulling images, starting containers, waiting for them to become available,
// setting up their initial state and cleaning up in the end.
//
// Its power is in a variety of Presets, each implementing a specific database,
// service or other tools. Each preset provides ways of setting up its initial
// state as easily as possible: SQL schema creation, test data upload into S3,
// sending test events to Splunk, etc.
//
// All containers created using Gnomock have a self-destruct mechanism that
// kicks-in right after the test execution completes.
//
// To debug cases where containers don't behave as expected, there are options
// like `WithDebugMode()` or `WithLogWriter()`.
//
// For the list of presets, please refer to
// https://pkg.go.dev/github.com/orlangure/gnomock/preset.
//
// Each preset can then be used in the following way:
//
//	p := redis.Preset() // replace "redis" with whatever you need
//	container, err := gnomock.Start(p)
//	addr := container.DefaultAddress() // e.g localhost:54321
package gnomock

import (
	"context"
	"io"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const defaultTag = "latest"

// newG creates a new Gnomock session with a unique identifier and a dedicated
// logger. It allows to follow a specific action while having multiple
// operations running in parallel.
func newG(debug bool) (*g, error) { _ = "STUB: not implemented"; return nil, nil }

// g is a Gnomock operation wrapper, mostly for debug purposes.
type g struct {
	id  uuid.UUID
	log *zap.SugaredLogger
}

// StartCustom creates a new container using provided image and binds random
// ports on the host to the provided ports inside the container. Image may
// include tag, which is set to "latest" by default. Optional configuration is
// available through Option functions. The returned container must be stopped
// when no longer needed using its Stop() method.
func StartCustom(image string, ports NamedPorts, opts ...Option) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newContainer(g *g, image string, ports NamedPorts, config *Options) (c *Container, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyf(dst io.Writer, src io.Reader) func() error { _ = "STUB: not implemented"; return nil }

func closeLogReader(logReader io.ReadCloser, g *errgroup.Group) func() error {
	_ = "STUB: not implemented"
	return nil
}

// Start creates a container using the provided Preset. The Preset provides its
// own Options to configure Gnomock container. Usually this is enough, but it
// is still possible to extend/override Preset options with new values. For
// example, wait timeout defined in the Preset, if at all, might be not enough
// for this particular usage, so it can't be changed during this call.
//
// All provided Options are applied. First, Preset options are applied. Then,
// custom Options. If both Preset and custom Options change the same
// configuration, custom Options are used.
//
// It is recommended, but not required, to call `gnomock.Stop()` when the tests
// complete to cleanup the containers.
func Start(p Preset, opts ...Option) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Stop stops the provided container and lets docker remove them from the
// system. Stop returns an error if any one of the containers couldn't stop. If
// these containers have sidecar containers, they are stopped as well.
func Stop(cs ...*Container) error { _ = "STUB: not implemented"; return nil }

func (g *g) stop(c *Container) error { _ = "STUB: not implemented"; return nil }

// Stop the sidecar container (best-effort) before returning, even if
// stopping the main container fails. Sidecar stop errors are intentionally
// ignored because sidecar containers have a self-destruct timer and will
// be cleaned up automatically.

func buildImage(image string) string { _ = "STUB: not implemented"; return "" }

func (g *g) setupLogForwarding(c *Container, cli *docker, config *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *g) wait(ctx context.Context, c *Container, config *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *g) initf(ctx context.Context, c *Container, config *Options) error {
	_ = "STUB: not implemented"
	return nil
}

// envAwareClone returns a copy of the provided container adjusted for usage
// inside current environment. For example, if current process runs directly on
// the host where container ports are exposed, an exact copy will be returned.
// For a process running itself inside a container, Host value will be replaced
// by docker host IP address. Anyway, calling Address() on the returned
// container will allow to communicate with it both from inside another
// container or from docker host.
func envAwareClone(c *Container) *Container { _ = "STUB: not implemented"; return nil }

// when gnomock runs inside docker container, the other container is only
// accessible through the host

func isHostDockerInternalAvailable() bool { _ = "STUB: not implemented"; return false }
