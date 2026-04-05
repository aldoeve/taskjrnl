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
	isNewDB := !isDBExists()
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
	CREATE TABLE IF NOT EXISTS Tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		tag TEXT,
		date_created TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		priority CHAR(1) NOT NULL CHECK(priority IN('L', 'M','H')),
		importance_variance INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS Positions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		position INTEGER NOT NULL,
		FOREIGN KEY (task_id) REFERENCES Tasks(id)
	)
	`
	_, err := db.Exec(schema)

	return err
}
