package store

import (
	"database/sql"
	"taskjrnl/internal/store/queries"
)

// Makes sure the requests positions exist.
func AreTasksInDB(db *sql.DB, positionA, positionB int) error {
	stmt := queries.SelectTaskIdGivenPositionSQL
	var taskId int
	err := db.QueryRow(stmt, positionA).Scan(&taskId)
	if err != nil {
		return err
	}
	return db.QueryRow(stmt, positionB).Scan(&taskId)
}
