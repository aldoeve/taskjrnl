package store

import (
	"database/sql"
	"taskjrnl/internal/store/queries"
)

// Removes a task and its refrences.
func RemoveTask(db *sql.DB, position_id int) error {
	stmt := queries.SelectTaskIdGivenPositionSQL
	var task_id int
	err := db.QueryRow(stmt, position_id).Scan(&task_id)
	if err != nil {
		return err
	}

	stmt = queries.DeletePositionRowGivenTaskIdSQL
	_, err = db.Exec(stmt, task_id)
	if err != nil {
		return err
	}

	stmt = queries.SelectPageIDsFromTaskIdSQL
	pageIds, err := db.Query(stmt, task_id)
	if err != nil {
		return err
	}
	defer pageIds.Close()

	for pageIds.Next() {
		var pageId int
		if err := pageIds.Scan(&pageId); err != nil {
			return err
		}

		stmt = queries.DeletePageFromPageIdSQL
		_, err = db.Exec(stmt, pageId)
		if err != nil {
			return nil
		}

	}

	stmt = queries.DeleteJrnlsFromTaskIdSQL
	_, err = db.Exec(stmt, task_id)
	if err != nil {
		return err
	}

	stmt = queries.DeleteTaskGivenTaskIdSQL
	_, err = db.Exec(stmt, task_id)
	if err != nil {
		return err
	}

	RearangePositions(db)

	return nil
}
