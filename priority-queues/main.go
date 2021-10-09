package main

import (
	"errors"
	"fmt"
)

type PriorityQueue struct {
	High []int
	Low  []int
}

// adds element to the end of our queue
func (q *PriorityQueue) Enqueue(elem int, highPriority bool) {
	if highPriority {
		q.High = append(q.High, elem)
	} else {
		q.Low = append(q.Low, elem)
	}
}

// returns the first element of the queue
func (q *PriorityQueue) Dequeue() (int, error) {

	// check high priorirty queue is not empty
	if len(q.High) != 0 {
		// get first element frim high priorirty queue
		firstElement := q.High[0]

		// update high queue to start at position 1 (removing index 0)
		q.High = q.High[1:]

		return firstElement, nil
	}

	// check low priorirty queue is not empty
	if len(q.Low) != 0 {
		// get first element frim low priorirty queue
		firstElement := q.Low[0]

		// update low queue to start at position 1 (removing index 0)
		q.Low = q.Low[1:]

		return firstElement, nil
	}

	return 0, errors.New("empty queue")

}

func (q *PriorityQueue) Length() int {
	return len(q.High) + len(q.Low)
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.Length() == 0
}

// returns first element of queue without updating queue
func (q *PriorityQueue) Peek() (int, error) {
	if len(q.High) != 0 {
		return q.High[0], nil
	}

	if len(q.Low) != 0 {
		return q.Low[0], nil
	}

	return 0, errors.New("empty queue")
}

func main() {
	// Priority Queue is similar to the basic queue
	// but it cotains a high priority queue and a low one
	queue := PriorityQueue{}
	fmt.Println(queue) // {[] []}

	// first test case
	queue.Enqueue(1, true)
	fmt.Println(queue) // {[1] []}

	queue.Enqueue(2, true)
	fmt.Println(queue) // {[1 2][]}

	elem, _ := queue.Dequeue()
	fmt.Println(elem)  // 1
	fmt.Println(queue) // {[2] []}

	fmt.Println("--------")

	// second test case
	queue.Enqueue(3, false)
	fmt.Println(queue) // {[2] [3]}

	queue.Enqueue(4, false)
	fmt.Println(queue) // {[2] [3 4]}

	elem, _ = queue.Dequeue()
	fmt.Println(elem)  // 2
	fmt.Println(queue) // {[] [3 4]}
}
