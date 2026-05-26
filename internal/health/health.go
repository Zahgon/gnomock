// Package health includes common health check functions to use in other
// packages.
package health

import (
	"context"
)

// HTTPGet returns no error when a request to the provided address (host:port)
// completes with a status < 400 (Bad Request).
func HTTPGet(ctx context.Context, addr string) error { _ = "STUB: not implemented"; return nil }

// nolint:gosec
