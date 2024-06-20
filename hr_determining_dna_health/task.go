package task

const INF = 1_000_000_000_000_000

type Strand struct {
	Left, Right int
	Value       string
}

type Node struct {
	Children map[string]*Node
	Values   [][]int
}

// Insert the value into a trie
func (n *Node) Insert(subStr string, index, value int) {
	if len(subStr) == 1 {
		if n.Children == nil {
			n.Children = make(map[string]*Node)
		}

		// check if transition via char is present already in children
		for char, child := range n.Children {
			if char == subStr[0:1] {
				if child.Values == nil {
					child.Values = [][]int{
						{index, value},
					}

					return
				}

				child.Values = append(child.Values, []int{index, value})

				return
			}
		}

		// no such child, so inserting a new one
		newNode := &Node{Values: [][]int{
			{index, value},
		}}

		n.Children[subStr] = newNode

		return
	}

	if n.Children == nil {
		n.Children = make(map[string]*Node)
	}

	for char, child := range n.Children {
		if char == subStr[0:1] {
			child.Insert(subStr[1:], index, value)

			return
		}
	}

	newNode := &Node{Values: [][]int{}}
	// didn't find an existing children with provided letter
	n.Children[subStr[0:1]] = newNode
	newNode.Insert(subStr[1:], index, value)
}

func (n *Node) Score(left, right int, strand string, index int) int {
	if n == nil {
		return 0
	}

	score := 0
	// indices are sorted as we inserted them in ascending order
	for i := 0; i < len(n.Values); i++ {
		if n.Values[i][0] > right {
			break
		}

		if n.Values[i][0] >= left {
			score += n.Values[i][1]
		}
	}

	if index < len(strand) {
		if _, ok := n.Children[string(strand[index])]; ok {
			score += n.Children[string(strand[index])].Score(left, right, strand, index+1)
		}
	}

	return score
}

func solve(n, s int, genes []string, healths []int, strands []Strand) (int, int) {
	if n == 0 {
		return 0, 0
	}

	root := Node{}
	for i := 0; i < n; i++ {
		root.Insert(genes[i], i, healths[i])
	}

	minSum := INF
	maxSum := 0
	for _, strand := range strands {
		totalSum := 0
		for i := 0; i < len(strand.Value); i++ {
			totalSum += root.Score(strand.Left, strand.Right, strand.Value, i)
		}

		if minSum > totalSum {
			minSum = totalSum
		}
		if maxSum < totalSum {
			maxSum = totalSum
		}
	}

	return minSum, maxSum
}
