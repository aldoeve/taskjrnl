package store

import "database/sql"

func RemoveTask(db *sql.DB, position_id int) error {
	findTaskIdFromPosition := `
		SELECT T.id
		FROM Tasks T
		RIGHT JOIN Positions P
		ON T.id = P.task_id
		WHERE P.Position = ?;
	`
	var task_id int
	err := db.QueryRow(findTaskIdFromPosition, position_id).Scan(&task_id)
	if err != nil {
		return err
	}

	deleteTask := `
		DELETE FROM Tasks
		WHERE id = ?
	`
	_, err = db.Exec(deleteTask)
	if err != nil {
		return err
	}

	RearangePositions(db)

	return nil
}
