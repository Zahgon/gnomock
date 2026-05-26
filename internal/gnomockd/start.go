package gnomockd

import (
	"io"
	"net/http"

	"github.com/orlangure/gnomock"
)

func startHandler() http.HandlerFunc { _ = "STUB: not implemented"; return *new(http.HandlerFunc) }

func setupLogWriter(done chan bool) (io.Writer, chan []string) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil
}

type startRequest struct {
	Options gnomock.Options `json:"options"`
	Preset  gnomock.Preset  `json:"preset"`
}
