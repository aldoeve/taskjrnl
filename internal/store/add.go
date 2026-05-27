package store

import (
	"database/sql"
	consts "taskjrnl/internal/consts"
	schema "taskjrnl/internal/schema"
	"taskjrnl/internal/store/queries"

	_ "modernc.org/sqlite"
)

// Inserts task into the database.
func insertTask(db *sql.DB, task schema.Tasks) error {
	const stmt = queries.InsertSingleTaskSQL

	_, err := db.Exec(stmt, task.Name,
		task.Tag, task.Priority,
		consts.DefaultVairance)

	if err != nil {
		return err
	}

	return RearangePositions(db)
}

// Creates a single task and saves it. Pointer parameters are optional so nil is equivalent to choosing defaults.
func CreateTask(db *sql.DB, taskName string, tag *string, priority *string) error {
	finalPriority := consts.LowPriority

	if priority != nil {
		finalPriority = *priority
	}

	task := schema.Tasks{
		Name:     taskName,
		Tag:      tag,
		Priority: &finalPriority,
	}

	return insertTask(db, task)
}
