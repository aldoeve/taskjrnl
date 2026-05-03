package util_test

import (
	"container/heap"
	"taskjrnl/internal/schema"
	util "taskjrnl/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pq_insertionNdeletion(t *testing.T) {
	pq := &util.PositionPriorityQueue{}
	heap.Init(pq)

	priority := "L"

	firstTask := schema.Tasks{Id: 45, ImportanceVariance: 10, Priority: &priority}
	secondTask := schema.Tasks{Id: 17, ImportanceVariance: 20, Priority: &priority}

	assert.Equal(t, 0, pq.Len())
	heap.Push(pq, &firstTask)
	assert.Equal(t, 1, pq.Len())
	heap.Push(pq, &secondTask)
	assert.Equal(t, 2, pq.Len())

	popped := heap.Pop(pq)
	assert.Equal(t, 1, pq.Len())
	assert.Equal(t, 17, popped)
	popped = heap.Pop(pq)
	assert.Equal(t, 0, pq.Len())
	assert.Equal(t, 45, popped)
}

func Test_pq_bubbleup(t *testing.T) {
	pq := &util.PositionPriorityQueue{}
	heap.Init(pq)

	priority := "L"

	firstTask := schema.Tasks{Id: 45, ImportanceVariance: 20, Priority: &priority}
	secondTask := schema.Tasks{Id: 17, ImportanceVariance: 10, Priority: &priority}

	heap.Push(pq, &firstTask)
	heap.Push(pq, &secondTask)

	assert.Equal(t, 2, pq.Len())
	top := heap.Pop(pq)
	assert.Equal(t, 45, top)
}
