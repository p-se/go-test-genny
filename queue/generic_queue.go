package queue

import (
	"container/list"

	"github.com/cheekybits/genny/generic"
)

// Generic is the type of the queue.
type Generic generic.Type

// GenericQueue is the type of the queue.
type GenericQueue struct {
	l *list.List
}

// NewGenericQueue creates a new queue.
func NewGenericQueue() GenericQueue {
	return GenericQueue{list.New()}
}

// Enqueue adds an item at the end of the queue.
func (q *GenericQueue) Enqueue(item Generic) {
	q.l.PushBack(item)
}

// Dequeue removes and returns the item on top of the queue.
func (q *GenericQueue) Dequeue() Generic {
	listElement := q.l.Front()
	q.l.Remove(listElement)
	value := listElement.Value
	return value.(Generic)
}

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Generic=BUILTINS"
