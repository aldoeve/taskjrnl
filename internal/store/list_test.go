package store_test

import (
	"path/filepath"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fetchOneTask(t *testing.T) {
	db, err := store.DBconnection(filepath.Join(t.TempDir(), consts.TestDBName))
	assert.Nil(t, err)

	tasks, err := store.FetchAllTasks(db)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(tasks))

	err = store.CreateTask(db, "Yep", nil, nil)
	assert.Nil(t, err)

	tasks, err = store.FetchAllTasks(db)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tasks))

	assert.Nil(t, db.Close())
}
