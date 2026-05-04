// The package errors contains custom errors made for debugging this application.
package taskjrnlErrors

import "errors"

var (
	ErrTooFewArgs   = errors.New("Too few arguments given")
	IncorrectFormat = errors.New("Incorrect format for requested mode")
)
