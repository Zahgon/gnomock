// Package errors exposes errors types used by gnomockd endpoints.
package errors

import (
	"github.com/orlangure/gnomock"
)

// NewPresetNotFoundError is returned when an invalid/unknown preset name was
// used.
func NewPresetNotFoundError(name string) error { _ = "STUB: not implemented"; return nil }

type presetNotFoundError struct {
	name   string
	ErrStr string `json:"error"`
}

func (e presetNotFoundError) Error() string {
	_ = "STUB: not implemented"

	// NewInvalidStartRequestError means that the request parameters of /start call
	// were invalid.
	return ""
}

func NewInvalidStartRequestError(err error) error { _ = "STUB: not implemented"; return nil }

type invalidStartRequestError struct {
	err    error
	ErrStr string `json:"error"`
}

func (e invalidStartRequestError) Error() string {
	_ = "STUB: not implemented"

	// NewStartFailedError means that the container failed to start for some
	// reason.
	return ""
}

func NewStartFailedError(err error, c *gnomock.Container) error {
	_ = "STUB: not implemented"
	return nil
}

type startFailedError struct {
	err       error
	ErrStr    string             `json:"error"`
	Container *gnomock.Container `json:"container,omitempty"`
}

func (e startFailedError) Error() string {
	_ = "STUB: not implemented"

	// InvalidStopRequestError means that the request parameters of /stop call were
	// invalid.
	return ""
}

func InvalidStopRequestError(err error) error { _ = "STUB: not implemented"; return nil }

type invalidStopRequestError struct {
	err    error
	ErrStr string `json:"error"`
}

func (e invalidStopRequestError) Error() string {
	_ = "STUB: not implemented"

	// StopFailedError means that the container failed to stop.
	return ""
}

func StopFailedError(err error, c *gnomock.Container) error { _ = "STUB: not implemented"; return nil }

type stopFailedError struct {
	err       error
	ErrStr    string             `json:"error"`
	Container *gnomock.Container `json:"container,omitempty"`
}

func (e stopFailedError) Error() string {
	_ = "STUB: not implemented"

	// ErrorCode returns HTTP response code for the provided error.
	return ""
}

func ErrorCode(err error) int { _ = "STUB: not implemented"; return 0 }
