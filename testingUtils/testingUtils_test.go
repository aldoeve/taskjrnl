package testingutils_test

import (
	"fmt"
	testingutils "taskjrnl/testingUtils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_captureOutput(t *testing.T) {
	const expectedOutput = "Hello"

	actualOutput := testingutils.CaptureOutput(t,
		func() { fmt.Print(expectedOutput) },
	)

	assert.Equal(t, expectedOutput, actualOutput)
}
