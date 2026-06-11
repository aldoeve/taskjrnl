package store

import (
	"database/sql"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/schema"
	"taskjrnl/internal/store/queries"
)

// Updates the required task in the database.
func updateTask(db *sql.DB, task schema.Tasks) error {
	const stmt = queries.UpdateTaskAllColumnsSQL

	_, err := db.Exec(stmt, task.Name, task.Tag, task.Priority, task.Id)
	if err != nil {
		return err
	}

	if task.Priority != nil {
		return RearangePositions(db)
	}

	return err
}

// Organized information to prepare to update task.
func ModifyTask(db *sql.DB, userViewTaskId int, name string, tag *string, priority *string) error {
	finalPriority := consts.LowPriority

	if priority != nil {
		finalPriority = *priority
	}

	stmt := queries.SelectTaskIdGivenPositionSQL
	var taskId int
	err := db.QueryRow(stmt, userViewTaskId).Scan(&taskId)
	if err != nil {
		return err
	}

	task := schema.Tasks{
		Id:       taskId,
		Name:     name,
		Tag:      tag,
		Priority: &finalPriority,
	}

	return updateTask(db, task)
}
