package dsalgo

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	currPath := make([]int, 0)
	currSum := 0
	pathTraversal(root, targetSum, currPath, currSum, &result)
	return result
}

func pathTraversal(root *TreeNode, targetSum int, currPath []int, currSum int, result *[][]int) {
	if root == nil {
		return
	}
	currPath = append(currPath, root.Val)
	currSum = currSum + root.Val
	if targetSum == currSum && root.Right == nil && root.Left == nil {
		*result = append(*result, currPath)
		return
	}
	if root.Left != nil {
		pathTraversal(root.Left, targetSum, append(make([]int, 0), currPath...), currSum, result)
	}
	if root.Right != nil {
		pathTraversal(root.Right, targetSum, append(make([]int, 0), currPath...), currSum, result)

	}
}
