package queue4

import (
	"errors"
	"hometask2/node"
)

type Queue4[T comparable] struct {
	Front  *node.Node[T]
	Rear   *node.Node[T]
	Length int
}

func NewQueue[T comparable]() *Queue4[T] {
	return &Queue4[T]{
		Front:  nil,
		Rear:   nil,
		Length: 0,
	}
}

func (q *Queue4[T]) Enqueue(value T) {
	newNode := &node.Node[T]{
		Data: value,
		Next: nil,
	}

	if q.Rear == nil {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}
	q.Length++
}

func (q *Queue4[T]) Dequeue() (T, error) {
	var zero T

	if q.Front == nil {
		return zero, errors.New("очередь пуста")
	}

	value := q.Front.Data
	q.Front = q.Front.Next

	if q.Front == nil {
		q.Rear = nil
	}

	q.Length--
	return value, nil
}

func (q *Queue4[T]) Peek() (T, error) {
	var zero T
	if q.Front == nil {
		return zero, errors.New("очередь пуста")
	}
	return q.Front.Data, nil
}

func (q *Queue4[T]) IsEmpty() bool {
	return q.Front == nil
}

func (q *Queue4[T]) Clear() {
	q.Front = nil
	q.Rear = nil
	q.Length = 0
}

func (q *Queue4[T]) Size() int {
	return q.Length
}
