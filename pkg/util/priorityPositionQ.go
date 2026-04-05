package util

type PositonItem struct {
	task_id  int
	priority int
	index    int
}
type PositonPriorityQueue []*PositonItem

func (pq PositonPriorityQueue) Len() int { return len(pq) }

func (pq PositonPriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}
