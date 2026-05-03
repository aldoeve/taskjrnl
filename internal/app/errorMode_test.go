package app_test

import (
	"taskjrnl/internal/app"
	"testing"

	errors "taskjrnl/internal/errors"

	"github.com/stretchr/testify/assert"
)

func Test_errorReturn(t *testing.T) {
	err := app.NoCorrespondingMode(nil)
	assert.Equal(t, err, errors.ErrUsage)
}
