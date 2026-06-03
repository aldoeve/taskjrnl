package store

import (
	"database/sql"
	"taskjrnl/internal/schema"
	"taskjrnl/internal/store/queries"
)

func FetchTaskNotes(db *sql.DB, taskId int) ([]schema.Pages, error) {
	return []schema.Pages{}, nil
}

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

	if err != nil {
		return task, err
	}

	return task, nil

}
