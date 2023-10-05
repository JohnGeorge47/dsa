package dsalgo

//https://leetcode.com/problems/binary-tree-right-side-view/

/*
Me thinks this is bfs so thing is at each level we will only be able to see
the rightmost element so we will look at each level now in case of root its ok

What we need to keep in mind is when we push left n right from node to the queue
from root (base)
The len(queue) is basically count on all element at that level in tree
so if we loop through the queue till the len(queue) then we will get all elements
at that level
We also add elements to the queue but right first then left  and so the rightmost element
or the i=0 at a particular level will be view at the level
You may get a doubt like wait arent we modifying the queue like at every level by
doing an append but we remember we run the loop based on stored qval and only once
that is over do we proceed to the next level
*/
func rightSideView(root *TreeNode) []int {
	bfsq := make([]*TreeNode, 0)
	rightView := make([]int, 0)
	if root == nil {
		return rightView
	}
	bfsq = append(bfsq, root)
	for len(bfsq) != 0 {
		sizeOfq := len(bfsq)
		for i := 0; i < sizeOfq; i++ {
			if i == 0 {
				rightView = append(rightView, bfsq[0].Val)
			}
			node := bfsq[0]
			bfsq = bfsq[1:]
			if node.Right != nil {
				bfsq = append(bfsq, node.Right)
			}
			if node.Left != nil {
				bfsq = append(bfsq, node.Left)
			}
		}
	}
	return rightView
}
