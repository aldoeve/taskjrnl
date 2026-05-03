package util

import (
	schema "taskjrnl/internal/schema"
)

type PositonItem struct {
	taskId     int
	importance int
	index      int
}
type PositionPriorityQueue []*PositonItem

func (pq PositionPriorityQueue) Len() int { return len(pq) }

func (pq PositionPriorityQueue) Less(i, j int) bool {
	return pq[i].importance > pq[j].importance
}

func (pq PositionPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Pushes a task id from a given task onto the priority queue.
func (pq *PositionPriorityQueue) Push(x any) {
	task := x.(*schema.Tasks)

	item := PositonItem{
		taskId:     task.Id,
		importance: CalculateImportance(task),
		index:      len(*pq),
	}

	*pq = append(*pq, &item)
}

// Returns taskId with the highest priority.
func (pq *PositionPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	positonItem := old[n-1]
	old[n-1] = nil
	positonItem.index = -1
	*pq = old[0 : n-1]
	return positonItem.taskId
}

/* func (pq *PositionPriorityQueue) update(positionItem *PositonItem, taskId int, importance int) {
	positionItem.taskId = taskId
	positionItem.importance = importance
	heap.Fix(pq, positionItem.index)
} */
