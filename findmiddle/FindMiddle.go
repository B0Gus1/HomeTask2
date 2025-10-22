package findmiddle

import (
	"hometask2/node"
)

func FindMiddle(head *node.Node[int]) *node.Node[int] {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
