package app

import (
	"database/sql"
	"taskjrnl/internal/store"
)

func ListMode(db *sql.DB) error {
	return store.ListAllTasks(db)
}
