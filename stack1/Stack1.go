package stack1

import (
	"errors"
	"hometask2/node"
)

type Stack1[T comparable] struct {
	Top  *node.Node[T]
	Size int
}

func NewStack[T comparable]() *Stack1[T] {
	return &Stack1[T]{Top: nil, Size: 0}
}

func (s *Stack1[T]) Push(value T) {
	newNode := &node.Node[T]{
		Data: value,
		Next: s.Top,
	}
	s.Top = newNode
	s.Size++
}

func (s *Stack1[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("стек пуст")
	}

	value := s.Top.Data
	s.Top = s.Top.Next
	s.Size--

	return value, nil
}

func (s *Stack1[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("стек пуст")
	}
	return s.Top.Data, nil
}

func (s *Stack1[T]) IsEmpty() bool {
	return s.Top == nil
}
