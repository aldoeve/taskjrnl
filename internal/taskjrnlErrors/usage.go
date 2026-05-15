// Package taskjrnlErrors holds custom errors that the application can raise.
package taskjrnlErrors

import "errors"

var (
	// Generic error message for user error.
	ErrUsage = errors.New("Incorrect Usage")
)
