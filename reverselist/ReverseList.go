package reverselist

import (
	"hometask2/node"
)

func ReverseList(head *node.Node[int]) *node.Node[int] {
	var prev *node.Node[int]
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	head = prev
	return head

}
