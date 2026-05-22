package store_test

import (
	"math/rand/v2"
	"path/filepath"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fetchNone(t *testing.T) {
	db, err := store.DBconnection(filepath.Join(t.TempDir(), consts.TestDBName))
	assert.Nil(t, err)

	tasks, err := store.FetchAllTasks(db)
	assert.Nil(t, err)

	assert.Equal(t, 0, len(tasks))
	assert.Nil(t, db.Close())
}
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

func Test_fetchAllTasks(t *testing.T) {
	db, err := store.DBconnection(filepath.Join(t.TempDir(), consts.TestDBName))
	assert.Nil(t, err)

	phrases := []string{"THE", "TREE", "IS", "GROWING"}
	const numOfTasks = 50

	for i := 0; i < numOfTasks; i++ {
		randomIndex := rand.IntN(len(phrases))
		secondRandomIndex := rand.IntN(len(phrases))

		store.CreateTask(db, phrases[randomIndex]+phrases[secondRandomIndex], nil, nil)
	}

	tasks, err := store.FetchAllTasks(db)
	assert.Nil(t, err)

	assert.Equal(t, numOfTasks, len(tasks))
	assert.Nil(t, db.Close())
}
