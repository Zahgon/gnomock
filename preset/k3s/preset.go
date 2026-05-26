// Package k3s provides a Gnomock Preset for lightweight kubernetes (k3s). This
// preset by no means should be used in any kind of deployment, and no other
// presets are supposed to be deployed in it. The goal of this preset is to
// allow easier testing of Kubernetes automation tools.
//
// This preset uses the `docker.io/rancher/k3s` image on Docker Hub as described
// by the [K3s documentation](https://docs.k3s.io/advanced#running-k3s-in-docker.)
//
// > ```bash
// > $ docker run \
// > --privileged \
// > --name k3s-server-1 \
// > --hostname k3s-server-1 \
// > -p 6443:6443 \
// > -d rancher/k3s:v1.24.10-k3s1 \
// > server
// > ```
//
// Please make sure to pick a version here:
// https://hub.docker.com/r/rancher/k3s/tags.
//
// Keep in mind that k3s runs in a single docker container, meaning it might be
// limited in memory, CPU and storage. Also remember that this cluster always
// runs on a single node.
//
// To connect to this cluster, use `Config` function that can be used together
// with Kubernetes client for Go, or `ConfigBytes` that can be saved as
// `kubeconfig` file and used by `kubectl`.
package k3s

import (
	"context"
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/internal/registry"
	"k8s.io/client-go/rest"
)

const (
	// defaultAPIPort is the default port that the K3s HTTPS Kubernetes API gets
	// served over.
	defaultAPIPort = 48443
	// defaultVersion is the default k3s version to run.
	defaultVersion = "v1.26.3-k3s1"
)

const (
	// KubeconfigPort is a port that exposes a single `/kubeconfig.yaml`
	// endpoint. It can be used to retrieve a configured kubeconfig file to use
	// to connect to this container using kubectl.
	kubeconfigPort = 48480

	// KubeConfigPortName is the name of the kubeconfig port that serves the
	// `kubeconfig.yaml`.
	KubeConfigPortName = "kubeconfig"

	// k3sManifestsDir is the directory with the K3s container where manifests
	// get automatically applied from.
	k3sManifestsDir = "/var/lib/rancher/k3s/server/manifests/"
)

// kubeconfigHttpd is a representation of the httpd manifest for k3s to
// automatically apply that will serve the k3s admin kubeconfig at
// `/kubeconfig.yaml`.
var kubeconfigHttpd = map[string]interface{}{
	"apiVersion": "v1",
	"kind":       "Pod",
	"metadata": map[string]interface{}{
		"name":      "kubeconfig-httpd",
		"namespace": "kube-system",
	},
	"spec": map[string]interface{}{
		"hostNetwork": true,
		"containers": []map[string]interface{}{
			{
				"name":  "web",
				"image": "docker.io/library/busybox:latest",
				"command": []string{
					"httpd", "-f", "-v",
					"-p", strconv.Itoa(kubeconfigPort),
				},
				"workingDir": "/var/gnomock/",
				"ports": []map[string]interface{}{
					{
						"name":          "http",
						"containerPort": kubeconfigPort,
						"protocol":      "TCP",
					},
				},
				"volumeMounts": []map[string]interface{}{
					{
						"name":      "kubeconfig-dir",
						"mountPath": "/var/gnomock/",
					},
				},
			},
		},
		"volumes": []map[string]interface{}{
			{
				"name": "kubeconfig-dir",
				"hostPath": map[string]interface{}{
					"path": "/var/gnomock/",
					"type": "Directory",
				},
			},
		},
	},
}

// kubeConfigHTTPJSONBytes is a representation of kubeconfigHttpd as a JSON
// encoded byte-array.
var kubeConfigHTTPJSONBytes []byte

// reServerAddress is a compiled regular expression that matches on the K3s
// API address to replace.
var reServerAddress *regexp.Regexp

func init() {
	registry.Register("kubernetes", func() gnomock.Preset { return &P{} })

	kubeConfigHTTPJSONBytesLocal, err := json.Marshal(kubeconfigHttpd)
	if err != nil {
		panic(err)
	}

	kubeConfigHTTPJSONBytes = kubeConfigHTTPJSONBytesLocal

	reServerAddress = regexp.MustCompile(`https://127.0.0.1:\d+`)
}

// Preset creates a new Gmomock k3s preset. This preset includes a
// k3s specific healthcheck function and default k3s image and port. Please
// note that this preset launches a privileged docker container.
//
// By default, this preset sets up k3s v1.19.3.
func Preset(opts ...Option) gnomock.Preset { _ = "STUB: not implemented"; return *new(gnomock.Preset) }

// P is a Gnomock Preset implementation of lightweight kubernetes (k3s).
type P struct {
	Version string `json:"version"`
	// Port is the API port for K3s to listen on.
	Port int

	// UseDynamicPort instructs the preset to use a dynamic host port instead of
	// a static one.
	UseDynamicPort bool

	// K3sServerFlags are additional k3s server flags added by options.
	K3sServerFlags []string
}

// Image returns an image that should be pulled to create this container.
func (p *P) Image() string { _ = "STUB: not implemented"; return "" }

// Ports returns ports that should be used to access this container.
func (p *P) Ports() gnomock.NamedPorts { _ = "STUB: not implemented"; return *new(gnomock.NamedPorts) }

// Options returns a list of options to configure this container.
func (p *P) Options() []gnomock.Option { _ = "STUB: not implemented"; return nil }

func (p *P) healthcheck(ctx context.Context, c *gnomock.Container) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// this is valid only for health checks, and solves a problem where
// gnomockd performs these calls from within its own container by accessing
// the cluster at 172.0.0.1, which is not one of the addresses in the
// certificate

func (p *P) setDefaults() { _ = "STUB: not implemented"; return }

// ConfigBytes returns file contents of kubeconfig file that should be used to
// connect to the cluster running in the provided container.
func ConfigBytes(c *gnomock.Container) (configBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Config returns `*rest.Config` instance of Kubernetes client-go package. This
// config can be used to create a new client that will work against k3s cluster
// running in the provided container.
func Config(c *gnomock.Container) (*rest.Config, error) { _ = "STUB: not implemented"; return nil, nil }
