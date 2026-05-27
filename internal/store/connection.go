// Package store holds the applications storage functions and core database.
package store

import (
	"database/sql"
	"os"
	consts "taskjrnl/internal/consts"
	"taskjrnl/internal/store/queries"

	_ "modernc.org/sqlite"
)

// Returns true on database file existing.
func isDBExists(dbLocation string) bool {
	_, err := os.Stat(dbLocation)
	return err == nil
}

// Returns a connection to the database.
func DBconnection(dbLocation string) (*sql.DB, error) {
	isNewDB := !isDBExists(dbLocation)

	db, err := sql.Open(consts.DatabaseType, dbLocation)
	if isNewDB && err == nil {
		err = initSchema(db)
	}

	if err == nil {
		_, _ = db.Exec(queries.JournalModeWAL)
		_, err = db.Exec(queries.ForeignKeysON)
	}

	return db, err
}

// Creates the database's tables.
func initSchema(db *sql.DB) error {
	_, err := db.Exec(queries.SchemaSQL)
	return err
}
