package node

type Node[T comparable] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}
