package removeelement

import (
	"hometask2/node"
)

func RemoveElement(head *node.Node[int], val int) *node.Node[int] {
	dummy := &node.Node[int]{}
	dummy.Next = head
	prev := dummy
	cur := head

	for cur != nil {
		if cur.Data == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return dummy.Next

}
