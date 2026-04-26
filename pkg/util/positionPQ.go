package util

import (
	"container/heap"
	schema "taskjrnl/internal/schema"
)

type PositonItem struct {
	taskId     int
	importance int
	index      int
}
type PositonPriorityQueue []*PositonItem

func (pq PositonPriorityQueue) Len() int { return len(pq) }

func (pq PositonPriorityQueue) Less(i, j int) bool {
	return pq[i].importance > pq[j].importance
}

func (pq PositonPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

// Pushes a task id from a given task onto the priority queue.
func (pq *PositonPriorityQueue) Push(x any) {
	task := x.(*schema.Tasks)

	item := PositonItem{
		taskId:     task.Id,
		importance: CalculateImportance(task),
		index:      len(*pq),
	}

	*pq = append(*pq, &item)
}

// Returns taskId with the highest priority.
func (pq *PositonPriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	positonItem := old[n-1]
	old[n-1] = nil
	positonItem.index = -1
	*pq = old[0 : n-1]
	return positonItem.taskId
}

func (pq *PositonPriorityQueue) update(positionItem *PositonItem, taskId int, importance int) {
	positionItem.taskId = taskId
	positionItem.importance = importance
	heap.Fix(pq, positionItem.index)
}
