package tool

import (
	"github.com/evenfound/even-go/node/cmd/evec/config"

	"github.com/pkg/errors"
	"github.com/ztrue/tracerr"
)

// Must be error-free, panic otherwise.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Ignore skips an error.
func Ignore(err error) {
	if err != nil {
	}
}

// NewError creates a new error.
func NewError(message string) error {
	if !config.Debug {
		return errors.New(message)
	}
	return tracerr.Wrap(errors.New(message))
}

// Wrap adds stacktrace and context message to an error.
// Stack trace causes a performance overhead, so used in debug mode only.
func Wrap(err error, message string) error {
	if !config.Debug {
		return errors.WithMessage(err, message)
	}
	return tracerr.Wrap(errors.WithMessage(err, message))
}

// Wrapf adds stacktrace and formatted context message to an error.
// Stack trace causes a performance overhead, so used in debug mode only.
func Wrapf(err error, format string, args ...interface{}) error {
	if !config.Debug {
		return errors.WithMessagef(err, format, args...)
	}
	return tracerr.Wrap(errors.WithMessagef(err, format, args...))
}
