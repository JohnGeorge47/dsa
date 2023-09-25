package dsalgo

//Problem:https://leetcode.com/problems/min-stack/description/

/*
I donno I felt here I should implement a stack from scratch which will do all
the operations though i do see an implementation using stdlib stack
but anyways here I decided to use a linked list and the stack having the head

The important thing to note is the structure of the node each node has a min element along
with usual next and val
*/

type MinStack struct {
	Head *node
}

type node struct {
	val  int
	min  int
	next *node
}

func Constructors() MinStack {
	return MinStack{Head: nil}
}

/*
The whole magic happens in the push where we check two things
 1. If the stack is empty push the first value by initializing the head with val and set min as val
    since min was empty
 2. If the stack is not we change the head to a new one and point the next to old one
    then we do a min check
*/
func (this *MinStack) Push(val int) {
	if this.Head == nil {
		this.Head = &node{
			val:  val,
			min:  val,
			next: nil,
		}
	} else {
		currhead := this.Head
		this.Head = &node{
			val: val,
			/*
				Now this is where we need to pay attention coz this makes everything work
				so what we need to know is every next head will have the previous min.
				so whenever something new is pushed and it is min the "new head will only have
				this min" the next head will have the previous min which can be literally anywhere behind
				as long as we dont pop that particular node
			*/
			min:  min(currhead.min, val),
			next: currhead,
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	this.Head = this.Head.next
}

func (this *MinStack) Top() int {
	return this.Head.val
}

func (this *MinStack) GetMin() int {
	return this.Head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
