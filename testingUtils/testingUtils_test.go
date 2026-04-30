package testingutils_test

import (
	"fmt"
	"os"
	"os/exec"
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

func Test_changingAndRestoringDir(t *testing.T) {
	expected, err := os.Getwd()
	assert.Nil(t, err)

	originalPerFunc := testingutils.ChangeWD(t, "../")
	testingutils.RestoreWD(t, originalPerFunc)

	actual, err := os.Getwd()
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func Test_removeDB(t *testing.T) {
	originalPath := testingutils.ChangeWD(t, "../")
	defer testingutils.RestoreWD(t, originalPath)

	cmd := exec.Command("touch", "./internal/store/TJ.db")
	cmd.Run()

	testingutils.RemoveDB(t)
	_, err := os.Stat("./internal/store/TJ.db")

	assert.NotNil(t, err)
}
