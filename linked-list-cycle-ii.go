package dsalgo

// https://leetcode.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	start := head
	for start != fast {
		start = start.Next
		fast = fast.Next
	}
	return fast
}
