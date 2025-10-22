package queue2

import "errors"

type Queue2[T any] struct {
	Items    []T
	Front    int
	Rear     int
	Size     int
	Capacity int
}

func NewQueue[T any](capacity int) *Queue2[T] {
	return &Queue2[T]{
		Items:    make([]T, capacity),
		Front:    0,
		Rear:     -1,
		Size:     0,
		Capacity: capacity,
	}
}

func (q *Queue2[T]) Enqueue(item T) error {
	if q.IsFull() {
		return errors.New("очередь переполнена")
	}

	q.Rear = (q.Rear + 1) % q.Capacity
	q.Items[q.Rear] = item
	q.Size++

	return nil
}

func (q *Queue2[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("очередь пуста")
	}

	item := q.Items[q.Front]
	q.Front = (q.Front + 1) % q.Capacity
	q.Size--

	return item, nil
}

func (q *Queue2[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("очередь пуста")
	}

	return q.Items[q.Front], nil
}

func (q *Queue2[T]) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue2[T]) IsFull() bool {
	return q.Size == q.Capacity
}

func (q *Queue2[T]) Clear() {
	q.Front = 0
	q.Rear = -1
	q.Size = 0
}
