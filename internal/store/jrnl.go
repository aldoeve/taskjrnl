package store

import (
	"database/sql"
	"taskjrnl/internal/store/queries"
)

func AddNoteToTask(db *sql.DB, position_id int, note string) error {
	stmt := queries.SelectTaskIdGivenPositionSQL
	var task_id int
	err := db.QueryRow(stmt, position_id).Scan(&task_id)
	if err != nil {
		return err
	}

	stmt = queries.CreatePageSQL
	result, err := db.Exec(stmt, note)
	if err != nil {
		return err
	}

	page_id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	stmt = queries.CreateJrnlSQL
	_, err = db.Exec(stmt, task_id, page_id)
	if err != nil {
		return err
	}

	return nil
}
