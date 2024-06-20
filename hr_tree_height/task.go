package task

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("err")
		return
	}

	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numbers := strings.Split(input, " ")

	nodes := make([]int, 0, len(numbers))

	for _, num := range numbers {
		parsedNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("err")
			return
		}

		nodes = append(nodes, parsedNum)
	}

	solve(n, nodes)
}
