package dsalgo

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil {
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		val := l1Val + l2Val + carry
		carry = val / 10
		val = val % 10
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}
