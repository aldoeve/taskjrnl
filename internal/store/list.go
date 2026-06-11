package store

import (
	"database/sql"
	schema "taskjrnl/internal/schema"
	"taskjrnl/internal/store/queries"

	_ "modernc.org/sqlite"
)

// Returns array containing every task ordered by most important first.
func FetchAllTasks(db *sql.DB) ([]schema.Tasks, error) {
	var tasks []schema.Tasks

	const stmt = queries.SelectRelavantOrderedListInfoSQL
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
