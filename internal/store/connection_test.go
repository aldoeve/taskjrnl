package store_test

import (
	"path/filepath"
	"taskjrnl/internal/consts"
	store "taskjrnl/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_establishConnToNewDB(t *testing.T) {
	tempDir := t.TempDir()
	dbLocation := filepath.Join(tempDir, consts.TestDBName)

	conn, err := store.DBconnection(dbLocation)

	assert.Nil(t, err)

	err = conn.Close()

	assert.Nil(t, err)
}

func Test_establishConnToOldDB(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	query := `
		INSERT INTO Tasks (name, tag, priority, importance_variance)
		VALUES ('TEST_OLD_DB', 'x', 'L', 5);
	`
	_, err = db.Exec(query)
	assert.Nil(t, err)
	assert.Nil(t, db.Close())

	db, err = store.DBconnection(dbLocation)
	assert.Nil(t, err)

	query = `
		SELECT COUNT(*) FROM Tasks WHERE name='TEST_OLD_DB';
	`
	row := db.QueryRow(query)
	var numberOfTasks int
	err = row.Scan(&numberOfTasks)
	assert.Nil(t, err)

	assert.Equal(t, numberOfTasks, 1)
	assert.Nil(t, db.Close())
}

func Test_schemaIsCorrect(t *testing.T) {
	dbLocation := filepath.Join(t.TempDir(), consts.TestDBName)

	db, err := store.DBconnection(dbLocation)
	assert.Nil(t, err)

	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table';`)
	assert.Nil(t, err)

	var tables []string
	for rows.Next() {
		var tableName string
		assert.Nil(t, rows.Scan(&tableName))
		tables = append(tables, tableName)
	}

	assert.Contains(t, tables, "Tasks")
	assert.Contains(t, tables, "Positions")
	assert.Nil(t, db.Close())
	assert.Nil(t, rows.Close())
}
