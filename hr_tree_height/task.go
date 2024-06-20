package task

import (
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(node *Node) {
	if n.Value > node.Value {
		if n.Left == nil {
			n.Left = node

			return
		}

		n.Left.Insert(node)

		return
	}

	if n.Right == nil {
		n.Right = node

		return
	}

	n.Right.Insert(node)
}

func (n *Node) Height() int {
	if n == nil {
		return -1
	}

	left := n.Left.Height()
	right := n.Right.Height()

	return int(math.Max(float64(left), float64(right))) + 1
}

func solve(n int, nodes []int) int {
	if n == 0 {
		return 0
	}

	root := Node{Value: nodes[0]}
	for i := 1; i < n; i++ {
		node := &Node{Value: nodes[i]}
		root.Insert(node)
	}

	return root.Height()
}
