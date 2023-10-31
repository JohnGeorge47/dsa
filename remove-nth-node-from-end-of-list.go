package dsalgo

import "fmt"

//https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/*
When in doubt use dummy node
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	headcpy := head
	i := 0
	count := 0
	for headcpy != nil {
		headcpy = headcpy.Next
		count++
	}
	for head != nil {
		fmt.Println(head.Val)
		if i == count-n {

		} else {
			curr.Next = &ListNode{Val: head.Val}
			curr = curr.Next
		}
		head = head.Next
		i++
	}
	return dummy.Next
}

/*
Fast and slow pointer solution also exists
*/
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
