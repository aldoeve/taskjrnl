package main

import (
	"math/rand/v2"
	"os"
	"taskjrnl/internal/consts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkForMinArgs_noArgs(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{consts.AppName}

	err := checkForMinArgs()

	assert.NotNil(t, err)
}

func Test_checkForMinArgs_1arg(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{consts.AppName, "1_arg"}

	err := checkForMinArgs()

	assert.Nil(t, err)
}

func Test_checkForMinArgs_multipleArgs(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	const fillerArgs = "Yep"
	os.Args = []string{consts.AppName, "1_arg", "2_arg"}
	numOfArgsToAdd := rand.IntN(11)

	for range numOfArgsToAdd {
		os.Args = append(os.Args, fillerArgs)
	}

	err := checkForMinArgs()

	assert.Nil(t, err)
}

func Test_positionialHelpArg(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{consts.AppName, "help"}

}
