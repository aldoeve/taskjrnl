package store_test

import (
	"fmt"
	"os/exec"
	"taskjrnl/internal/consts"
	store "taskjrnl/internal/store"
	testingutils "taskjrnl/testingUtils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RemoveDB(t *testing.T) error {
	cmd := exec.Command("rm", consts.DBLocation)
	out, err := cmd.CombinedOutput()
	assert.Nil(t, err)
	fmt.Print(out)
	return nil
}

func Test_establishConnToNewDB(t *testing.T) {
	originalPath := testingutils.ChangeWD(t, "../..")
	defer testingutils.RestoreWD(t, originalPath)
	defer RemoveDB(t)

	conn, err := store.DBconnection()
	assert.Nil(t, err)
	err = conn.Close()
	assert.Nil(t, err)
}
