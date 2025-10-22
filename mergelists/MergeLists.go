package mergelists

import "hometask2/node"

func MergeLists(head1, head2 *node.Node[int]) *node.Node[int] {
	dummy := &node.Node[int]{}
	current := dummy

	for head1 != nil && head2 != nil {
		if head1.Data <= head2.Data {
			current.Next = head1
			head1 = head1.Next
		} else {
			current.Next = head2
			head2 = head2.Next
		}
		current = current.Next
	}

	if head1 != nil {
		current.Next = head1
	} else {
		current.Next = head2
	}

	return dummy.Next
}
