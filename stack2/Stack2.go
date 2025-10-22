package stack2

import "errors"

type Stack2[T any] struct {
	Items []T
	Top   int
	Size  int
}

func NewStack[T any](size int) *Stack2[T] {
	return &Stack2[T]{
		Items: make([]T, size),
		Top:   -1,
		Size:  size,
	}
}

func (s *Stack2[T]) Push(item T) error {
	if s.IsFull() {
		return errors.New("стек полон")
	}

	s.Top++
	s.Items[s.Top] = item
	return nil
}

func (s *Stack2[T]) Pop() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, errors.New("стек пуст")
	}

	item := s.Items[s.Top]
	s.Top--
	return item, nil
}

func (s *Stack2[T]) Peek() (T, error) {
	var zero T

	if s.IsEmpty() {
		return zero, errors.New("стек пуст")
	}

	return s.Items[s.Top], nil
}

func (s *Stack2[T]) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack2[T]) IsFull() bool {
	return s.Top == s.Size-1
}

func (s *Stack2[T]) Clear() {
	s.Top = -1
}
