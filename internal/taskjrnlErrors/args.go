package taskjrnlErrors

import "errors"

var (
	// Error message to indicate incorrect number of arguments.
	ErrTooFewArgs = errors.New("Too few arguments given")

	// Error message to indicate a format error.
	IncorrectFormat = errors.New("Incorrect format for requested mode")
)
