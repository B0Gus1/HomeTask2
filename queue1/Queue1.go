package queue1

import (
	"errors"
	"hometask2/node"
)

type Queue1[T comparable] struct {
	Front *node.Node[T]
	Rear  *node.Node[T]
	Size  int
}

func NewQueue[T comparable]() *Queue1[T] {
	return &Queue1[T]{
		Front: nil,
		Rear:  nil,
		Size:  0,
	}
}

func (q *Queue1[T]) Enqueue(data T) {
	newNode := &node.Node[T]{
		Data: data,
		Next: nil,
		Prev: q.Rear,
	}

	if q.Rear == nil {

		q.Front = newNode
		q.Rear = newNode
	} else {

		q.Rear.Next = newNode
		q.Rear = newNode
	}
	q.Size++
}

func (q *Queue1[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("очередь пуста")
	}

	data := q.Front.Data
	if q.Size == 1{
		q.Front = nil
		q.Rear = nil
	} else {
		q.Front = q.Front.Next
		q.Front.Prev = nil
	}

	q.Size--
	return data, nil
}

func (q *Queue1[T]) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue1[T]) Clear() {
	q.Front = nil
	q.Rear = nil
	q.Size = 0
}

func (q *Queue1[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("очередь пуста")
	}
	return q.Front.Data, nil
}
