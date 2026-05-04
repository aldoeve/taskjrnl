package app

import (
	"database/sql"
	"taskjrnl/internal/schema"
	"taskjrnl/internal/store"
)

func drawTasks(tasks []schema.Tasks) {
	return
}

func ListMode(db *sql.DB) error {
	tasks, err := store.ListAllTasks(db)
	if err != nil {
		return err
	}

	drawTasks(tasks)
	return nil
}
