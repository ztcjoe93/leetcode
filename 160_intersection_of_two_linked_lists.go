/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	sliceA := make([]*ListNode, 0)
	sliceB := make([]*ListNode, 0)

	for headA != nil {
		sliceA = append(sliceA, headA)
		headA = headA.Next
	}

	for headB != nil {
		sliceB = append(sliceB, headB)
		headB = headB.Next
	}

	curA := len(sliceA) - 1
	curB := len(sliceB) - 1

	var intersection *ListNode
	for curA >= 0 && curB >= 0 {
		if fmt.Sprintf("%p", sliceA[curA]) == fmt.Sprintf("%p", sliceB[curB]) {
			intersection = sliceA[curA]
			curA--
			curB--
		} else {
			break
		}
	}

	return intersection
}