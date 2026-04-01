package util

import "taskjrnl/internal/schema"

type PositionQItem struct {
	Positon schema.Positions
	value   int
	index   int
}

type PriorityQueue []*PositionQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].value < pq[j].value
}
