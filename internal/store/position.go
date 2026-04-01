package store

import (
	"database/sql"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func RearangePositions(db *sql.DB, task_id int64) error {
	stmt := `
	SELECT Positions.position, Tasks.id, 
	Tasks.date_created, Tasks.priority, 
	Tasks.importance_variance
	FROM Positions
	Join Tasks
	ON Positions.task_id = Tasks.id
	ORDER BY Positions.position
	`
	rows, err := db.Query(stmt)
	if err != nil {
		println(err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next() {
		positionInfo := schema.Positions{}
		taskInfo := schema.Tasks{}

		if err := rows.Scan(
			&positionInfo.Position,
			&taskInfo.Id,
			&taskInfo.DateCreated,
			&taskInfo.Priority,
			&taskInfo.ImportanceVariance,
		); err != nil {
			println(err.Error())
			return err
		}
		println(positionInfo.Position, taskInfo.Id)
	}

	return nil
}
