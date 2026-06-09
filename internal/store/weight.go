package store

import (
	"database/sql"
	"taskjrnl/internal/store/queries"
)

// Takes a new weight and adds it to what the db has for the given task.
func AdjustTaskWeight(db *sql.DB, userTaskPosition int, adjustment int) error {
	stmt := queries.SelectTaskIdGivenPositionSQL
	var taskId int
	if err := db.QueryRow(stmt, userTaskPosition).Scan(&taskId); err != nil {
		return err
	}

	stmt = queries.UpdateTaskWeightSQL
	_, err := db.Exec(stmt, adjustment, taskId)
	if err != nil {
		return err
	}

	return RearangePositions(db)
}
