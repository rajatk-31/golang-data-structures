package main

import "fmt"

type Queue struct {
	items []int
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}

// Enqueue
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// Dequeue
func (q *Queue) Dequeue() int {
	removed := q.items[0]
	q.items = q.items[1:]
	return removed
}
