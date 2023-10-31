package dsalgo

//https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/submissions/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	lvl := true
	for len(queue) != 0 {
		qlen := len(queue)
		lvlVals := make([]int, 0)
		if lvl {
			for i := 0; i < qlen; i++ {
				lvlVals = append(lvlVals, queue[i].Val)
			}
		} else {
			for i := qlen - 1; i >= 0; i-- {
				lvlVals = append(lvlVals, queue[i].Val)
			}
		}
		res = append(res, lvlVals)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:len(queue)]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		lvl = !lvl
	}
	return res
}
