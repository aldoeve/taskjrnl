package store

import (
	"database/sql"
	schema "taskjrnl/internal/schema"

	_ "modernc.org/sqlite"
)

func FetchAllTasks(db *sql.DB) ([]schema.Tasks, error) {
	var tasks []schema.Tasks

	stmt := `
		SELECT T.name, T.tag, T.date_created, T.priority 
		FROM Tasks AS T
		INNER JOIN Positions AS P
			ON T.id = P.task_id
		ORDER BY P.position;
	`
	rows, err := db.Query(stmt)
	if err != nil {
		return tasks, err
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
			return tasks, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}
