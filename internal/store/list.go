package store

import (
	"database/sql"
	"fmt"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func ListAllTasks(db *sql.DB) error {
	stmt := `
		SELECT name, tag, create_date, priority FROM Tasks;
	`
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		task := schema.Tasks{}

		if err := rows.Scan(
			&task.Name,
			&task.Tag,
			&task.DateCreated,
			&task.Priority,
		); err != nil {
			return err
		}
		fmt.Println(task.Name, task.Tag, task.DateCreated, task.Priority)
	}
	return nil
}
