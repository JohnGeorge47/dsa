package dsalgo

//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
/*
Kind of harder compared to the other one which is binary search tree
lets look at an example
      3
    /   \
   5     1
  / \    / \
6    2  0    8

We always start at the root now lets our p and q (5,1)
now there are three cases we need to see

1. p and q are not descendents of either so the root is the lowest ancestor of either
2. p is a descendent of q that means q is the lca
3. q is descendent of p that means p is lca

*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/*
		Lets first look at the base case
		if our root is null that means no use going here
		now if we find p or q we return the values
	*/
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//We traverse both sides
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//This means that either p or q is on the left tree and is a
	//either a descendent of the other or is in the other subtree
	//in previous recursive steps
	if left == nil {
		return right
	}
	//	This means that either p or q is on the right  tree and is a
	//	either a descendent of the other or is in the other subtree
	if right == nil {
		return left
	}
	//Now if both are not nil this means that the root will be the lowest ancestor
	return root
}
