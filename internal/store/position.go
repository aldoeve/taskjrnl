package store

import (
	"database/sql"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func RearangePositions(db *sql.DB, task_id int64) error {
	stmt := `
	SELECT task_id, position FROM Positions 
	ORDER BY position
	`
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		positionRow := schema.Positions
	}

	return nil
}
