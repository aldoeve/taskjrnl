package store

import (
	"database/sql"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func cleanPositionsTable(db *sql.DB) error {
	stmt := `
	DELETE FROM Positions;
	`

	if _, err := db.Exec(stmt); err != nil {
		return err
	}

	return nil
}

func RearangePositions(db *sql.DB, task_id int64) error {

	if err := cleanPositionsTable(db); err != nil {
		return err
	}

	stmt := `
	SELECT 
	id, date_created, 
	priority, importance_variance
	FROM Tasks;
	`
	rows, err := db.Query(stmt)
	if err != nil {
		println(err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next() {
		taskInfo := schema.Tasks{}

		if err := rows.Scan(
			&taskInfo.Id,
			&taskInfo.DateCreated,
			&taskInfo.Priority,
			&taskInfo.ImportanceVariance,
		); err != nil {
			println(err.Error())
			return err
		}
	}

	return nil
}
