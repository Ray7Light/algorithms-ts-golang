package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack{}

	stack.push(5)
	stack.push(7)
	stack.push(9)

	popAndCheck(t, stack, 9)
	check(t, stack.length, 2)

	stack.push(11)

	popAndCheck(t, stack, 11)
	popAndCheck(t, stack, 7)

	peeked, err := stack.peek()

	if err == nil {
		check(t, peeked, 5)
	} else {
		t.Errorf("%s", err)
	}

	popAndCheck(t, stack, 5)

	val, err := stack.pop()

	if err != nil && val != 0 {
		t.Errorf("expected to return an error")
	}

	stack.push(69)

	peeked, err = stack.peek()

	if err == nil {
		check(t, peeked, 69)
	} else {
		t.Errorf("%s", err)
	}

	check(t, stack.length, 1)
}

func popAndCheck(t testing.TB, s *Stack, want int) {
	t.Helper()
	got, err := s.pop()

	if err != nil {
		t.Errorf("%s", err)
	} else {
		check(t, got, want)
	}
}

func check(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
