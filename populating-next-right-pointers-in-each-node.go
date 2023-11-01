package dsalgo

//https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		currLen := len(queue)
		for i := 0; i < currLen; i++ {
			currNode := queue[0]
			if currNode == nil {
				continue
			}
			queue = queue[1:]
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
			if i != currLen-1 {
				currNode.Next = queue[0]
			} else {
				currNode.Next = nil
			}
		}
	}
	return root
}
