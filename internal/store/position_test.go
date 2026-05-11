// Package store_test tests the applications functions that interact with the database.
package store_test

import (
	"path/filepath"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_positionTablePopulated(t *testing.T) {
	dblocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dblocation)
	assert.Nil(t, err)

	tasksNames := []string{"Hello"}

	err = store.CreateTask(db, tasksNames[0], nil, nil)
	assert.Nil(t, err)

	stmt := `
		SELECT T.name
		FROM Tasks as T
		INNER JOIN Positions AS P
			ON T.id = P.task_id
		ORDER BY P.position;
	`
	rows, err := db.Query(stmt)
	assert.Nil(t, err)

	var (
		count    int
		taskName string
	)

	tasksNamesLength := len(tasksNames)

	for rows.Next() {
		rows.Scan(&taskName)
		assert.Equal(t, tasksNames[tasksNamesLength-1-count], taskName)
		count++
	}

	assert.Equal(t, 1, count)
	assert.Nil(t, rows.Close())
	assert.Nil(t, db.Close())
}

func Test_positionTablePopulatedAndReordered(t *testing.T) {
	dblocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dblocation)
	assert.Nil(t, err)

	tasksNames := []string{"Hello", "Hi"}

	err = store.CreateTask(db, tasksNames[0], nil, nil)
	assert.Nil(t, err)

	customPriority := consts.HighPriority
	err = store.CreateTask(db, tasksNames[1], nil, &customPriority)

	stmt := `
		SELECT T.name
		FROM Tasks as T
		INNER JOIN Positions AS P
			ON T.id = P.task_id
		ORDER BY P.position;
	`
	rows, err := db.Query(stmt)
	assert.Nil(t, err)

	var (
		count    int
		taskName string
	)

	tasksNamesLength := len(tasksNames)

	for rows.Next() {
		rows.Scan(&taskName)
		assert.Equal(t, tasksNames[tasksNamesLength-1-count], taskName)
		count++
	}

	assert.Equal(t, 2, count)
	assert.Nil(t, rows.Close())
	assert.Nil(t, db.Close())
}
