package dsalgo

//https://leetcode.com/problems/validate-binary-search-tree/description/

/*
On first glance it looks simple af lets see its like comparing
parent with child but the caveat is all elements on left should be less than all root
elements in that subtree and all elements on elements on right should be greater than all roots in that side
*/

import "math"

func isValidBST(root *TreeNode) bool {
	minval := math.MinInt
	maxval := math.MaxInt
	// Think this it is a binary tree in root coz -∞<root.val<∞
	return recurseThroughTree(minval, maxval, root)
}

func recurseThroughTree(lLimit int, rLimit int, node *TreeNode) bool {
	if node == nil {
		return true
	}
	/*
	   this is the check to see if its between the limits
	*/
	if !((node.Val > lLimit) && (node.Val < rLimit)) {
		return false
	}
	/*
	   This is the important stuff happens
	   1 . For every left value the right limit will be the value of the node above since it cannot exceed that
	   2.  For every right value the left limit is value of node above it since it cannot be below it
	   Now we need to recursively do it lets look at a tree instead
	                5 (-∞<val<∞)
	              /  \
	             4       6 (5<val<∞)
	                 /              \
	                3 5<root.val<6       7 (5<root.val<∞)
	*/
	if recurseThroughTree(lLimit, node.Val, node.Left) && recurseThroughTree(node.Val, rLimit, node.Right) {
		return true
	}
	return false
}
