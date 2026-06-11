package store

import (
	"database/sql"
	"taskjrnl/internal/schema"
	"taskjrnl/internal/store/queries"
)

// Fetches a tasks notes sorted by oldest first.
func FetchTaskNotes(db *sql.DB, taskId int) ([]schema.Pages, error) {
	stmt := queries.FetchJrnlPagesFromTaskIdSQL
	rows, err := db.Query(stmt, taskId)
	if err != nil {
		return []schema.Pages{}, err
	}
	defer rows.Close()

	var pages []schema.Pages
	for rows.Next() {
		var page schema.Pages

		if err := rows.Scan(&page.Note, &page.DateCreated); err != nil {
			return pages, err
		}
		pages = append(pages, page)
	}

	return pages, nil
}

// Returnst the appropriate task info for the info keyword.
func FetchTaskinfo(db *sql.DB, postionalId int) (schema.Tasks, error) {
	task := schema.Tasks{}

	stmt := queries.SelectTaskIdGivenPositionSQL
	err := db.QueryRow(stmt, postionalId).Scan(&task.Id)
	if err != nil {
		return task, err
	}

	stmt = queries.SelectTaskInfoGivenTaskIdSQL
	err = db.QueryRow(stmt, task.Id).Scan(
		&task.Name, &task.Tag,
		&task.DateCreated, &task.Priority,
		&task.ImportanceVariance,
	)

	return task, err
}
