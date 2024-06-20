package task

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

func solve(n int, nodes []int) []int {
	if n == 0 {
		return []int{}
	}

	root := &Node{Value: nodes[0]}
	for i := 1; i < n; i++ {
		node := &Node{Value: nodes[i]}
		root.Insert(node)
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	i := 0
	results := make([]int, 0, len(nodes))

	for i < len(queue) {
		results = append(results, queue[i].Value)

		if queue[i].Left != nil {
			queue = append(queue, queue[i].Left)
		}
		if queue[i].Right != nil {
			queue = append(queue, queue[i].Right)
		}

		i++
	}

	return results
}
