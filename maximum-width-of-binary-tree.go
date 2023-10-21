package dsalgo

/*
Basically every node in a tree then the left n right are 2i and 2i+1


I think the solution is not the best tho
*/

func widthOfBinaryTree(root *TreeNode) int {
	q := make([]*TreeNode, 0)
	nodeMap := make(map[*TreeNode]int)
	nodeMap[root] = 1
	result := 0
	q = append(q, root)
	for len(q) != 0 {
		currLen := len(q)
		lvlLeft := 0
		lvlRight := 0
		for i := 0; i < currLen; i++ {
			node := q[0]
			q = q[1:]
			if i == 0 {
				lvlLeft = nodeMap[node]
			}
			if i == currLen-1 {
				lvlRight = nodeMap[node]
			}
			if node.Left != nil {
				nodeMap[node.Left] = 2 * nodeMap[node]
				q = append(q, node.Left)
			}
			if node.Right != nil {
				nodeMap[node.Right] = 2*nodeMap[node] + 1
				q = append(q, node.Right)
			}
		}
		size := lvlRight - lvlLeft + 1
		result = max(size, result)
	}
	return result
}
