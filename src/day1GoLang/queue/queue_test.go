package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}

	queue.enqueue(5)
	queue.enqueue(7)
	queue.enqueue(9)
	dequeueAndCheck(t, &queue, 5)
	check(t, queue.length, 2)
	queue.enqueue(11)
	dequeueAndCheck(t, &queue, 7)
	dequeueAndCheck(t, &queue, 9)

	peeked, err := queue.peek()

	if err != nil {
		t.Errorf("%s", err)
	} else if peeked != 11 {
		t.Errorf("peeked %d, want %d", peeked, 11)
	}

	dequeueAndCheck(t, &queue, 11)
	check(t, queue.length, 0)

	val, err := queue.dequeue()

	if val != 0 && err != nil {
		t.Errorf("expected to return error when the queue is empty")
	}

	queue.enqueue(69)

	peeked, err = queue.peek()

	if err != nil {
		t.Errorf("%s", err)
	} else if peeked != 69 {
		t.Errorf("peeked %d, want %d", peeked, 69)
	}

	check(t, queue.length, 1)
}

func dequeueAndCheck(t testing.TB, q *Queue, want int) {
	t.Helper()

	got, err := q.dequeue()

	if err == nil {
		check(t, got, want)
	} else {
		t.Errorf("%s", err)
	}
}

func check(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
