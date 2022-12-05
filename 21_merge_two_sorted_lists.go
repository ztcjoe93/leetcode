package main

func test_21() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	Mainly to remind me of the thought process when tackling questions involving LinkedList
	and nodes.

	Process is to have a dummy, where we store the address of it while having the tail node
	equate to it.

	Head: 0000
	Tail (points at head): 0000

	Thereafter, the tail node then starts to propagate to the node in either of the linked list.
	Then we set the tail node as the next node; repeat until the end of the lists.

	Then we return the next node from the headPtr, which will return the new linkedlist.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	headPtr := &ListNode{}
	tail := headPtr

	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			break
		}

		if list2 == nil {
			tail.Next = list1
			break
		}

		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	return headPtr.Next
}
