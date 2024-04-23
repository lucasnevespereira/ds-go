package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
}

// Enqueue adds element to the end of our queue
func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

// Dequeue returns the first element of the queue
func (q *Queue) Dequeue() (int, error) {

	// validate queue as not empty
	if len(q.Elements) == 0 {
		return 0, errors.New("queue is empty")
	}

	// get first element
	firstElement := q.Elements[0]

	// update elements slice to start at position 1 (removing index 0)
	q.Elements = q.Elements[1:]

	return firstElement, nil
}

func (q *Queue) Length() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

// returns first element of queue without updating queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return q.Elements[0], nil
}

func main() {
	// Queue is like a waiting list (first in / first out) (last in/ last out).
	// Similar to stack but open in both ends.
	// One end is always used to insert data (enqueue)
	// and the other is used to remove data (dequeue)
	queue := Queue{}
	fmt.Println(queue) // {[]}

	queue.Enqueue(1)
	fmt.Println(queue) // {[1]}

	queue.Enqueue(2)
	fmt.Println(queue) // {[1 2]}

	elem, _ := queue.Dequeue()
	fmt.Println(elem)  // 1
	fmt.Println(queue) // {[2]}
}
