package deque

import (
	"errors"
	"hometask2/node"
)

type Deque[T comparable] struct {
	Head *node.Node[T]
	Tail *node.Node[T]
	Size int
}

func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (d *Deque[T]) PushFront(value T) {
	newNode := &node.Node[T]{
		Data: value,
		Prev: nil,
		Next: d.Head,
	}

	if d.Head != nil {
		d.Head.Prev = newNode
	} else {
		d.Tail = newNode
	}

	d.Head = newNode
	d.Size++
}

func (d *Deque[T]) PushBack(value T) {
	newNode := &node.Node[T]{
		Data: value,
		Prev: d.Tail,
		Next: nil,
	}

	if d.Tail != nil {
		d.Tail.Next = newNode
	} else {
		d.Head = newNode
	}

	d.Tail = newNode
	d.Size++
}

func (d *Deque[T]) PopFront() (T, error) {
	var zero T
	if d.Head == nil {
		return zero, errors.New("дек пуст")
	}

	value := d.Head.Data

	if d.Head.Next != nil {
		d.Head.Next.Prev = nil
	} else {
		d.Tail = nil
	}

	d.Head = d.Head.Next
	d.Size--

	return value, nil
}

func (d *Deque[T]) PopBack() (T, error) {
	var zero T
	if d.Tail == nil {
		return zero, errors.New("дек пуст")
	}

	value := d.Tail.Data

	if d.Tail.Prev != nil {
		d.Tail.Prev.Next = nil
	} else {
		d.Head = nil
	}

	d.Tail = d.Tail.Prev
	d.Size--

	return value, nil
}

func (d *Deque[T]) Front() (T, error) {
	var zero T
	if d.Head == nil {
		return zero, errors.New("дек пуст")
	}
	return d.Head.Data, nil
}

func (d *Deque[T]) Back() (T, error) {
	var zero T
	if d.Tail == nil {
		return zero, errors.New("дек пуст")
	}
	return d.Tail.Data, nil
}

func (d *Deque[T]) IsEmpty() bool {
	return d.Size == 0
}

func (d *Deque[T]) Clear() {
	d.Head = nil
	d.Tail = nil
	d.Size = 0
}
