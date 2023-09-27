package dsalgo

//https://leetcode.com/problems/reverse-linked-list/description/

/*
The classic one hated it becaue i always got it wrong
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Lets first try the iterative solution so the thing is lets look at an example

	1->2->3->4
*/
func reverseList(head *ListNode) *ListNode {
	//lets first create a dummy node which will store previous value this will be
	//our result
	var prevNode *ListNode
	prevNode = nil
	for head != nil && head.Next != nil {
		//lets store next value since next value needs to be exchanged
		next := head.Next
		/*This is the important part the current heads next will be previous node
		 lets see some examples
		1.initial head =1-> prevNode=nil
		head.next=prev so 1->nil
		and prev=1->nil
		2.head 2-> prev=1->nil
		  head.next=prev
		  prev=2->1->nil
		3.head =3-> prev=2->1->nil
		  prev=3->2->1->nil
		*/
		head.Next = prevNode
		prevNode = head
		head = next
	}
	return prevNode
}

func reverseListRecursive(head *ListNode) *ListNode {
	return recursive(head, nil)
}

func recursive(currNode *ListNode, prevNode *ListNode) *ListNode {
	if currNode == nil {
		return prevNode
	}
	next := currNode.Next
	currNode.Next = prevNode
	prevNode = currNode
	return recursive(next, prevNode)
}
