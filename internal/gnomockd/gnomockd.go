// Package gnomockd is an HTTP wrapper around Gnomock
package gnomockd

import (
	"net/http"
)

// Handler returns an HTTP handler ready to serve incoming connections.
func Handler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func respondWithError(w http.ResponseWriter, err error) { _ = "STUB: not implemented"; return }
