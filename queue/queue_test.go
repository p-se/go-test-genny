package queue

import (
	"testing"
)

func TestGenericQueue(t *testing.T) {
	q := NewGenericQueue()
	v1 := new(Generic)
	q.Enqueue(v1)
	v2 := new(Generic)
	q.Enqueue(v2)

	if value := q.Dequeue(); value != v1 {
		t.Fatalf("Wrong order")
	}
	if value := q.Dequeue(); value != v2 {
		t.Fatalf("Wrong order")
	}
}

func TestStringQueue(t *testing.T) {
	q := NewStringQueue()
	q.Enqueue("one")
	q.Enqueue("two")

	if v := q.Dequeue(); v != "one" {
		t.Fatalf("v1 is not one")
	}
	if v := q.Dequeue(); v != "two" {
		t.Fatalf("v1 is not two")
	}
}
