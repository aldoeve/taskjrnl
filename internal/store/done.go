package store

import (
	"database/sql"
	"taskjrnl/internal/store/queries"
)

func RemoveTask(db *sql.DB, position_id int) error {
	stmt := queries.SelectTaskIdGivenPositionSQL
	var task_id int
	err := db.QueryRow(stmt, position_id).Scan(&task_id)
	if err != nil {
		return err
	}

	stmt = queries.DeletePositionRowGivenTaskIdSQL
	_, err = db.Exec(stmt, task_id)
	if err != nil {
		return err
	}

	stmt = queries.DeleteTaskGivenTaskIdSQL
	_, err = db.Exec(stmt, task_id)
	if err != nil {
		return err
	}

	RearangePositions(db)

	return nil
}
