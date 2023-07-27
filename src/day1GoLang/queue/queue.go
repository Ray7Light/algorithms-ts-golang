package queue

import (
	"errors"
)

type node struct {
	next  *node
	value int
}

type Queue struct {
	length int
	head   *node
	tail   *node
}

func (q *Queue) enqueue(val int) {
	newNode := &node{value: val}

	q.length++

	if q.tail == nil {
		q.tail = newNode
		q.head = q.tail
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) dequeue() (int, error) {
	if q.head == nil {
		return 0, errors.New("the queue is empty") // Return a default value or handle this case differently based on your requirement.
	}

	q.length--

	head := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil // Reset tail when the queue becomes empty.
	}

	return head.value, nil
}

func (q *Queue) peek() (int, error) {
	if q.head == nil {
		return 0, errors.New("the queue is empty") // Return a default value or handle this case differently based on your requirement.
	}

	return q.head.value, nil
}
