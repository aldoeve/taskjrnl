// Package testingutils holds testing utility functions.
package testingutils

import (
	"bytes"
	"io"
	"os"
	"syscall"
	"testing"

	"github.com/acarl005/stripansi"

	"github.com/stretchr/testify/assert"
)

// Returns a string of the captured stdout.
func CaptureOutput[T any](t *testing.T, f func(T) error, arg T) string {
	t.Helper()

	r, w, err := os.Pipe()
	assert.Nil(t, err)

	oldFd, err := syscall.Dup(int(os.Stdout.Fd()))
	assert.Nil(t, err)

	err = syscall.Dup2(int(w.Fd()), int(os.Stdout.Fd()))
	assert.Nil(t, err)

	err = f(arg)
	assert.Nil(t, err)

	err = syscall.Dup2(oldFd, int(os.Stdout.Fd()))
	assert.Nil(t, err)
	err = syscall.Close(oldFd)
	assert.Nil(t, err)

	err = w.Close()
	assert.Nil(t, err)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	assert.Nil(t, err)

	return stripansi.Strip(buf.String())
}
