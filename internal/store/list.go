package store

import (
	"database/sql"
	"fmt"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func ListAllTasks(db *sql.DB) error {
	stmt := `
		SELECT name, tag, create_date, priority, importance_variance FROM Tasks;
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
			&task.Create_date,
			&task.Priority,
			&task.Importance_variance,
		); err != nil {
			return err
		}
		fmt.Println(task.Name, task.Tag, task.Create_date, task.Priority, task.Importance_variance)
		println("there done")
	}
	println("finsihed")
	return nil
}
