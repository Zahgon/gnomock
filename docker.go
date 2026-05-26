package gnomock

import (
	"context"
	"io"
	"regexp"
	"sync"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/network"
	"github.com/moby/moby/client"
	"go.uber.org/zap"
)

const (
	localhostAddr             = "127.0.0.1"
	defaultStopTimeoutSec     = 1
	duplicateContainerPattern = `Conflict. The container name "(?:.+?)" is already in use by container "(\w+)". You have to remove \(or rename\) that container to be able to reuse that name.` // nolint:lll
	dockerSockAddr            = "/var/run/docker.sock"
)

var duplicateContainerRegexp = regexp.MustCompile(duplicateContainerPattern)

type docker struct {
	client *client.Client
	log    *zap.SugaredLogger

	// This lock is used to protect docker client from concurrent connections
	// with version negotiation. As of this moment, there is a data race in
	// docker client when version negotiation is requested. This data race is
	// not dangerous, but it triggers race detector alarms, so it should be
	// avoided. Currently the client still has this issue, so this is an
	// attempt to fix it locally by preventing concurrent connection using the
	// same client (mostly when `Stop` is called with multiple containers).
	//
	// https://github.com/moby/moby/pull/42379
	lock sync.Mutex
}

func (g *g) dockerConnect() (*docker, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *docker) isExistingLocalImage(ctx context.Context, image string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *docker) pullImage(ctx context.Context, image string, cfg *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *docker) startContainer(ctx context.Context, image string, ports NamedPorts, cfg *Options) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *docker) setupContainerCleanup(id string, cfg *Options) (chan string, context.CancelFunc) {
	_ = "STUB: not implemented"
	return nil, *new(context.CancelFunc)
}

func (d *docker) prepareContainer(
	ctx context.Context,
	image string,
	ports NamedPorts,
	cfg *Options,
) (*client.ContainerCreateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *docker) waitForContainerNetwork(ctx context.Context, id string, ports NamedPorts) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *docker) exposedPorts(namedPorts NamedPorts) network.PortSet {
	_ = "STUB: not implemented"
	return *new(network.PortSet)
}

func (d *docker) portBindings(exposedPorts network.PortSet, ports NamedPorts) network.PortMap {
	_ = "STUB: not implemented"
	return *new(network.PortMap)
}

// for the container to be accessible from another container, it cannot
// listen on 127.0.0.1 as it will be accessed by gateway address (e.g
// 172.17.0.1), so its port should be exposed everywhere

func (d *docker) createContainer(
	ctx context.Context,
	image string,
	ports NamedPorts,
	cfg *Options,
) (*client.ContainerCreateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *docker) findReusableContainer(
	ctx context.Context,
	image string,
	ports NamedPorts,
	cfg *Options,
) (*Container, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (d *docker) boundNamedPorts(inspectResp container.InspectResponse, namedPorts NamedPorts) (NamedPorts, error) {
	_ = "STUB: not implemented"
	return *new(NamedPorts), nil
}

func (d *docker) readLogs(ctx context.Context, id string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (d *docker) stopContainer(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *docker) stopClient() error { _ = "STUB: not implemented"; return nil }

func (d *docker) removeContainer(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// hostAddr returns an address of a host that runs the containers. If
// DOCKER_HOST environment variable is not set, if its value is an invalid URL,
// or if it is a `unix:///` socket address, it returns local address.
func (d *docker) hostAddr() string { _ = "STUB: not implemented"; return "" }

func isDeletionAlreadyInProgessError(err error, id string) bool {
	_ = "STUB: not implemented"
	return false
}
