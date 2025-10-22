package stack3

import "errors"

type Stack3[T any] struct {
	Items []T
	Size  int
}

func NewStack[T any](initialCapacity int) *Stack3[T] {
	return &Stack3[T]{
		Items: make([]T, 0, initialCapacity),
		Size:  0,
	}
}

func (s *Stack3[T]) Push(item T) {
	if s.Size == len(s.Items) {
		s.Resize(max(1, len(s.Items)*2))
	}

	if s.Size < len(s.Items) {
		s.Items[s.Size] = item
	} else {
		s.Items = append(s.Items, item)
	}
	s.Size++
}

func (s *Stack3[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("стек пуст")
	}

	s.Size--
	item := s.Items[s.Size]

	if s.Size > 0 && s.Size == len(s.Items)/4 {
		s.Resize(len(s.Items) / 2)
	}

	return item, nil
}

func (s *Stack3[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("стек пуст")
	}
	return s.Items[s.Size-1], nil
}

func (s *Stack3[T]) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack3[T]) Resize(newCapacity int) {
	newItems := make([]T, newCapacity)
	copy(newItems, s.Items[:s.Size])
	s.Items = newItems
}

func (s *Stack3[T]) Clear() {
	s.Items = make([]T, 0, cap(s.Items))
	s.Size = 0
}

func (s *Stack3[T]) Capacity() int {
	return len(s.Items)
}
