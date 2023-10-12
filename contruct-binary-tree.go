package dsalgo

//https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/*
Try drawing this out its a recursion by splitting the two arrays
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	root := TreeNode{
		Val: preorder[0],
	}
	mid := 0
	for i, i2 := range inorder {
		if i2 == preorder[0] {
			mid = i
		}
	}
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return &root
}
