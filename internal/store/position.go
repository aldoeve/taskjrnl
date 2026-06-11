package store

import (
	"container/heap"
	"database/sql"
	schema "taskjrnl/internal/schema"
	"taskjrnl/internal/store/queries"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"

	_ "modernc.org/sqlite"
)

// Deletes every row in the Positions table.
func clearPositionsTable(db *sql.DB) error {
	const stmt = queries.DeletePositionsRowsSQL

	_, err := db.Exec(stmt)

	return err
}

// Inserts task_id and its corresponding position into the Positions table.
func insertTaskIntoPosition(db *sql.DB, positionItem schema.Positions) error {
	const stmt = queries.InsertSinglePositionRowSQL

	_, err := db.Exec(stmt, positionItem.TaskId, positionItem.Position)

	return err
}

// Fixes the Positions table's ordering.
func RearangePositions(db *sql.DB) error {

	if err := clearPositionsTable(db); err != nil {
		return err
	}

	const stmt = queries.SelectPositionsDataFromTasksSQL
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Insertion into max priority queue to sort tasks' importances.
	pq := &util.PositionPriorityQueue{}
	heap.Init(pq)

	for rows.Next() {
		taskInfo := schema.Tasks{}

		if err := rows.Scan(
			&taskInfo.Id,
			&taskInfo.DateCreated,
			&taskInfo.Priority,
			&taskInfo.Weight,
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
