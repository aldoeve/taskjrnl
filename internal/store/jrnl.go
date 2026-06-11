package store

import (
	"database/sql"
	"taskjrnl/internal/store/queries"
)

// Addes an entry to the page table and jrnl table to save the note.
func AddNoteToTask(db *sql.DB, positionId int, note string) error {
	stmt := queries.SelectTaskIdGivenPositionSQL
	var taskId int
	err := db.QueryRow(stmt, positionId).Scan(&taskId)
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
	_, err = db.Exec(stmt, taskId, page_id)

	return err
}
