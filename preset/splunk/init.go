package splunk

import (
	"bytes"
	"context"
	"fmt"

	"github.com/orlangure/gnomock"
)

var errConflict = fmt.Errorf("409: conflict")

// Event is a type used during Splunk initialization. Pass events to WithValues
// to ingest them into the container before the control over it is passed to
// the caller.
type Event struct {
	// Event is the actual log entry. Can be any format
	Event string `json:"event"`

	// Index is the name of index to ingest the log into. If the index does not
	// exist, it will be created
	Index string `json:"index"`

	// Source will be used as "source" value of this event in Splunk
	Source string `json:"source"`

	// SourceType will be used as "sourcetype" value of this event in Splunk
	SourceType string `json:"sourcetype"`

	// Time represents event timestamp in seconds, milliseconds or nanoseconds
	// (and maybe even in microseconds, whatever splunk recognizes)
	Time int64 `json:"time"`
}

func (p *P) initf() gnomock.InitFunc { _ = "STUB: not implemented"; return *new(gnomock.InitFunc) }

// Ingest adds the provided events to splunk container. Use the same password
// you provided in WithPassword. Send as many events as you like, this function
// only returns when all the events were indexed, or when the context is timed
// out.
func Ingest(ctx context.Context, c *gnomock.Container, adminPassword string, events ...Event) error {
	_ = "STUB: not implemented"
	return nil
}

func ingestEvents(ctx context.Context, c *gnomock.Container, adminPassword string, events []Event) error {
	_ = "STUB: not implemented"
	return nil
}

func issueToken(post postFunc, addr string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func indexRegistry(post postFunc, addr string) func(string) error {
	_ = "STUB: not implemented"
	return nil
}

type ingestFunc func(Event) error

func eventForwarder(ctx context.Context, c *gnomock.Container, adminPassword string) (ingestFunc, error) {
	_ = "STUB: not implemented"
	return *new(ingestFunc), nil
}

type countFunc func() (int, error)

func eventCounter(ctx context.Context, c *gnomock.Container, adminPassword string) countFunc {
	_ = "STUB: not implemented"
	return *new(countFunc)
}

type postFunc func(string, *bytes.Buffer) ([]byte, error)

func requestWithAuth(ctx context.Context, method, password string, isJSON bool) postFunc {
	_ = "STUB: not implemented"
	return *new(postFunc)
}

type splunkTokenResponse struct {
	Entry []struct {
		Content struct {
			Token string `json:"token"`
		} `json:"content"`
	} `json:"entry"`
}

type splunkSearchcResponse struct {
	Result map[string]interface{} `json:"result"`
}
