package testingutils

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CaptureOutput(t *testing.T, f func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	assert.Nil(t, err)

	os.Stdout = w

	f()

	err = w.Close()
	assert.Nil(t, err)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	assert.Nil(t, err)
	os.Stdout = old

	return buf.String()
}
