package testingutils

import (
	"bytes"
	"io"
	"os"
	"os/exec"
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

func ChangeWD(t *testing.T, path string) string {
	originalPath, err := os.Getwd()
	assert.Nil(t, err)

	err = os.Chdir(path)
	assert.Nil(t, err)

	return originalPath
}

func RestoreWD(t *testing.T, path string) {
	err := os.Chdir(path)
	assert.Nil(t, err)
}

func RemoveDB(t *testing.T) {
	cmd := exec.Command("make", "clean")
	_, err := cmd.CombinedOutput()
	assert.Nil(t, err)
}
