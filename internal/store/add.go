package store

import (
	"database/sql"
	consts "taskjrnl/internal/consts"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func insertTask(db *sql.DB, task schema.Tasks) error {
	stmt := `
		INSERT INTO Tasks (name, tag, priority, importance_variance)
		VALUES(?, ?, ?, ?);
	`
	_, err := db.Exec(stmt, task.Name, task.Tag, task.Priority, consts.DefaultVairance)

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

// Creates and inserts a single task into the database. Pointer parameters are optional so nil means to chose defaults.
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
