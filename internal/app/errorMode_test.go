package app_test

import (
	"taskjrnl/internal/app"
	"testing"

	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"

	"github.com/stretchr/testify/assert"
)

func Test_errorReturn(t *testing.T) {
	err := app.NoCorrespondingMode(nil)
	assert.Equal(t, err, taskjrnlErrors.ErrUsage)
}
