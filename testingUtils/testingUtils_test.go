package testingutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_captureOutput(t *testing.T) {
	const expectedOutput = "Hello"

	actualOutput := captureOutput(t,
		func() { fmt.Print(expectedOutput) },
	)

	assert.Equal(t, expectedOutput, actualOutput)
}
