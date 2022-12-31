/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	addr := make(map[string]bool, 0)

	for head != nil {
		ptr := fmt.Sprintf("%p", head)
		if _, ok := addr[ptr]; ok {
			return true
		}
		addr[ptr] = true
		head = head.Next
	}

	return false
}