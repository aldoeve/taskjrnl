package store

import (
	"container/heap"
	"database/sql"
	schema "taskjrnl/internal/schema"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	util "taskjrnl/pkg/util"

	_ "modernc.org/sqlite"
)

func clearPositionsTable(db *sql.DB) error {
	stmt := `
		DELETE FROM Positions;
	`

	if _, err := db.Exec(stmt); err != nil {
		return err
	}

	return nil
}

func insertTaskIntoPosition(db *sql.DB, positionItem schema.Positions) error {
	stmt := `
		INSERT INTO Positions (task_id, position)
		VALUES(?, ?); 
	`
	_, err := db.Exec(stmt, positionItem.TaskId, positionItem.Position)
	if err != nil {
		return err
	}

	return nil
}

func RearangePositions(db *sql.DB) error {

	if err := clearPositionsTable(db); err != nil {
		return err
	}

	stmt := `
		SELECT 
		id, date_created, 
		priority, importance_variance
		FROM Tasks;
	`
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	pq := &util.PositionPriorityQueue{}
	heap.Init(pq)

	for rows.Next() {
		taskInfo := schema.Tasks{}

		if err := rows.Scan(
			&taskInfo.Id,
			&taskInfo.DateCreated,
			&taskInfo.Priority,
			&taskInfo.ImportanceVariance,
		); err != nil {
			return err
		}

		heap.Push(pq, &taskInfo)
	}

	var position int
	for pq.Len() > 0 {
		position++

		taskId, ok := heap.Pop(pq).(int)
		if !ok {
			return taskjrnlErrors.HeapPanic
		}

		positionItem := schema.Positions{
			TaskId:   taskId,
			Position: position,
		}

		err = insertTaskIntoPosition(db, positionItem)
		if err != nil {
			return err
		}
	}

	return nil
}
