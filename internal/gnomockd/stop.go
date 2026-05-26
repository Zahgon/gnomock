package gnomockd

import (
	"net/http"
)

func stopHandler() http.HandlerFunc { _ = "STUB: not implemented"; return *new(http.HandlerFunc) }

type stopRequest struct {
	ID string `json:"id"`
}
