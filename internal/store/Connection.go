// Package store holds the applications storage functions and core database.
package store

import (
	"database/sql"
	"os"
	consts "taskjrnl/internal/consts"

	_ "modernc.org/sqlite"
)

func isDBExists() bool {
	_, err := os.Stat(consts.DBLocation)
	return err == nil
}

func DBconnection() (*sql.DB, error) {
	isNewDB := isDBExists()
	db, err := sql.Open(consts.DatabaseType, consts.DBLocation)
	if isNewDB && err == nil {
		err = initSchema(db)
	}
	if err == nil {
		_, _ = db.Exec("PRAGMA journal_mode=WAL;")
		_, err = db.Exec("PRAGMA foreign_keys=ON;")
	}

	return db, err
}

func initSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		create_date TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		priority CHAR(1) NOT NULL CHECK(priority IN('L', 'M','H')),
		importance INTEGER NOT NULL
	);
	`
	_, err := db.Exec(schema)

	return err
}
