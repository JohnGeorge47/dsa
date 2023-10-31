package dsalgo

/*
Convert it to graph with hash map structure
then have a hash for visited set
*/

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := make(map[*TreeNode][]*TreeNode)
	queue := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]struct{})
	queue = append(queue, root)
	if len(queue) != 0 {
		node := queue[0]
		queue = queue[1:len(queue)]
		if node.Left != nil {
			graph[node] = append(graph[node], node.Left)
			graph[node.Left] = append(graph[node.Left], node)
		}
		if node.Right != nil {
			graph[node] = append(graph[node], node.Right)
			graph[node.Right] = append(graph[node.Right], node)
		}
	}
	dist := 0
	res := make([]int, 0)
	var neighbourQueue []*TreeNode
	copy(graph[target], neighbourQueue)
	for len(neighbourQueue) != 0 {
		dist = dist + 1
		currLen := len(neighbourQueue)
		for i := 0; i < currLen; i++ {
			node := neighbourQueue[0]
			queue = neighbourQueue[1:len(queue)]
			if _, ok := visited[node]; !ok {
				visited[node] = struct{}{}
				neighbourQueue = append(neighbourQueue, graph[node]...)
				if dist == k {
					res = append(res, node.Val)
				}
			}
		}
	}
	return res
}

func findNode(root *TreeNode, k int) *TreeNode {
	if root.Val == k {
		return root
	}
	if root == nil {
		return nil
	}
	left := findNode(root.Left, k)
	right := findNode(root.Right, k)
	if right == nil {
		return left
	}
	return right
}
