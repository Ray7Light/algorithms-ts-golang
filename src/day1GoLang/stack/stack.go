package stack

import (
	"errors"
)

type node struct {
	prev  *node
	value int
}

type Stack struct {
	head   *node
	length int
}

func (s *Stack) push(val int) {
	newNode := &node{value: val}

	s.length++

	if s.head == nil {
		s.head = newNode
		return
	}

	newNode.prev = s.head
	s.head = newNode
}

func (s *Stack) pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("the stack is empty")
	}

	s.length--

	node := s.head
	s.head = node.prev
	node.prev = nil

	return node.value, nil
}

func (s *Stack) peek() (int, error) {
	if s.head == nil {
		return 0, errors.New("the stack is empty")
	}

	return s.head.value, nil
}
