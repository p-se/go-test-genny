package main

import (
	"fmt"

	"github.com/p-se/go-test-genny/queue"
)

func main() {
	q := queue.NewStringQueue()
	q.Enqueue("foo")
	fmt.Print(q.Dequeue())
}
