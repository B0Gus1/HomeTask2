package queue3

import "errors"

type Queue3[T any] struct {
	Items    []T
	Head     int
	Tail     int
	Size     int
	Capacity int
}

func NewQueue[T any](initialCapacity int) *Queue3[T] {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}

	return &Queue3[T]{
		Items:    make([]T, initialCapacity),
		Head:     0,
		Tail:     0,
		Size:     0,
		Capacity: initialCapacity,
	}
}

func (q *Queue3[T]) Enqueue(item T) {
	if q.Size == q.Capacity {
		q.Resize(q.Capacity * 2)
	}

	q.Items[q.Tail] = item
	q.Tail = (q.Tail + 1) % q.Capacity
	q.Size++
}

func (q *Queue3[T]) Dequeue() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, errors.New("очередь пуста")
	}

	item := q.Items[q.Head]
	q.Items[q.Head] = zero
	q.Head = (q.Head + 1) % q.Capacity
	q.Size--

	if q.Size > 0 && q.Size == q.Capacity/4 {
		q.Resize(q.Capacity / 2)
	}

	return item, nil
}

func (q *Queue3[T]) Peek() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, errors.New("очередь пуста")
	}

	return q.Items[q.Head], nil
}

func (q *Queue3[T]) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue3[T]) Resize(newCapacity int) {
	if newCapacity < 10 {
		newCapacity = 10
	}

	newItems := make([]T, newCapacity)

	for i := 0; i < q.Size; i++ {
		index := (q.Head + i) % q.Capacity
		newItems[i] = q.Items[index]
	}

	q.Items = newItems
	q.Head = 0
	q.Tail = q.Size
	q.Capacity = newCapacity
}

func (q *Queue3[T]) Clear() {
	q.Items = make([]T, 10)
	q.Head = 0
	q.Tail = 0
	q.Size = 0
	q.Capacity = 10
}
