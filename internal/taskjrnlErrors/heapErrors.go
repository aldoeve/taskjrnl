package taskjrnlErrors

import "errors"

var (
	// Generic error for priority queue panics.
	HeapPanic = errors.New("Heap panic")
)
