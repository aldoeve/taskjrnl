package store_test

import (
	"path/filepath"
	"taskjrnl/internal/config"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/schema"
	"taskjrnl/internal/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_createAndInsertTaskPlain(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	taskName := "Task without tag and no priority"
	var (
		prioirty *string
		tag      *string
	)

	err = store.CreateTask(db, taskName, tag, prioirty)
	assert.Nil(t, err)

	stmt := `
	SELECT * 
	FROM Tasks;
	`

	rows, err := db.Query(stmt)
	assert.Nil(t, err)

	var (
		count int
		task  schema.Tasks
	)
	for rows.Next() {
		count++

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Tag,
			&task.DateCreated,
			&task.Priority,
			&task.ImportanceVariance,
		)
		assert.Nil(t, err)
	}

	assert.Equal(t, 1, count)
	assert.Equal(t, taskName, task.Name)
	assert.Nil(t, task.Tag)
	assert.Equal(t, consts.LowPriority, *task.Priority)
	assert.Equal(t, consts.DefaultVairance, task.ImportanceVariance)
	assert.Equal(t, time.Now().UTC().Format(config.TimeFormat), task.DateCreated)
	assert.Nil(t, rows.Close())
	assert.Nil(t, db.Close())
}

func Test_createAndInsertTaskTag(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	taskName := "Task without tag and no priority"
	var (
		prioirty *string
		tag      *string
	)
	simpleTag := "YEP"
	tag = &simpleTag

	err = store.CreateTask(db, taskName, tag, prioirty)
	assert.Nil(t, err)

	stmt := `
	SELECT * 
	FROM Tasks;
	`

	rows, err := db.Query(stmt)
	assert.Nil(t, err)

	var (
		count int
		task  schema.Tasks
	)
	for rows.Next() {
		count++

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Tag,
			&task.DateCreated,
			&task.Priority,
			&task.ImportanceVariance,
		)
		assert.Nil(t, err)
	}

	assert.Equal(t, 1, count)
	assert.Equal(t, taskName, task.Name)
	assert.Equal(t, *tag, *task.Tag)
	assert.Equal(t, consts.LowPriority, *task.Priority)
	assert.Equal(t, consts.DefaultVairance, task.ImportanceVariance)
	assert.Equal(t, time.Now().UTC().Format(config.TimeFormat), task.DateCreated)
	assert.Nil(t, rows.Close())
	assert.Nil(t, db.Close())
}

func Test_createAndInsertTaskPriority(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	taskName := "Task without tag and no priority"
	var (
		prioirty *string
		tag      *string
	)

	simplePriority := "H"
	prioirty = &simplePriority

	err = store.CreateTask(db, taskName, tag, prioirty)
	assert.Nil(t, err)

	stmt := `
	SELECT * 
	FROM Tasks;
	`

	rows, err := db.Query(stmt)
	assert.Nil(t, err)

	var (
		count int
		task  schema.Tasks
	)
	for rows.Next() {
		count++

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Tag,
			&task.DateCreated,
			&task.Priority,
			&task.ImportanceVariance,
		)
		assert.Nil(t, err)
	}

	assert.Equal(t, 1, count)
	assert.Equal(t, taskName, task.Name)
	assert.Nil(t, task.Tag)
	assert.Equal(t, consts.HighPriority, *task.Priority)
	assert.Equal(t, consts.DefaultVairance, task.ImportanceVariance)
	assert.Equal(t, time.Now().UTC().Format(config.TimeFormat), task.DateCreated)
	assert.Nil(t, rows.Close())
	assert.Nil(t, db.Close())
}

func Test_createAndInsertTaskFull(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	taskName := "Task without tag and no priority"
	var (
		prioirty *string
		tag      *string
	)
	simpleTag := "SKIL"
	simplePriority := "M"

	prioirty = &simplePriority
	tag = &simpleTag

	err = store.CreateTask(db, taskName, tag, prioirty)
	assert.Nil(t, err)

	stmt := `
	SELECT * 
	FROM Tasks;
	`

	rows, err := db.Query(stmt)
	assert.Nil(t, err)

	var (
		count int
		task  schema.Tasks
	)
	for rows.Next() {
		count++

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Tag,
			&task.DateCreated,
			&task.Priority,
			&task.ImportanceVariance,
		)
		assert.Nil(t, err)
	}

	assert.Equal(t, 1, count)
	assert.Equal(t, taskName, task.Name)
	assert.Equal(t, simpleTag, *task.Tag)
	assert.Equal(t, consts.MidPriority, *task.Priority)
	assert.Equal(t, consts.DefaultVairance, task.ImportanceVariance)
	assert.Equal(t, time.Now().UTC().Format(config.TimeFormat), task.DateCreated)
	assert.Nil(t, rows.Close())
	assert.Nil(t, db.Close())
}

func Test_createAndInsertTaskError(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	taskName := "Task without tag and no priority"
	var (
		prioirty *string
		tag      *string
	)

	errPriority := "INCORRECT"
	prioirty = &errPriority

	err = store.CreateTask(db, taskName, tag, prioirty)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "CHECK constraint failed: priority IN('L', 'M','H')")

	assert.Nil(t, db.Close())
}
