package store

import (
	"database/sql"
	consts "taskjrnl/internal/consts"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func insertTask(db *sql.DB, task schema.Tasks) error {
	stmt := `
		INSERT INTO Tasks (name, priority, importance_variance)
		VALUES(?, ?, ?);
	`
	_, err := db.Exec(stmt, task.Name, *task.Priority, *task.Importance_variance)

	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

// Creates and inserts a single task into the database. Pointer parameters are optional.
func CreateTask(db *sql.DB, taskName string, priority *string, importanceVariance *int) error {
	finalPriority := consts.LowPriority
	finalImportanceVariance := consts.InitalImportanceVariance

	if priority != nil {
		finalPriority = *priority
	}
	if importanceVariance != nil {
		finalImportanceVariance = *importanceVariance
	}

	task := schema.Tasks{
		Name:                taskName,
		Priority:            &finalPriority,
		Importance_variance: &finalImportanceVariance,
	}

	return insertTask(db, task)
}
