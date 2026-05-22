// Package testingutils_test tests the testing utilities.
package testingutils_test

import (
	"database/sql"
	"fmt"
	testingutils "taskjrnl/testingUtils"
	"testing"

	"github.com/stretchr/testify/assert"
)

const expectedOutput = "Hello"

func simpleOut(_ *sql.DB) error {
	fmt.Print(expectedOutput)
	return nil
}
func Test_captureOutput(t *testing.T) {

	actualOutput := testingutils.CaptureOutput(t, simpleOut, nil)

	assert.Equal(t, expectedOutput, actualOutput)
}
