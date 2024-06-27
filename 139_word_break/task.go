package task

type Trie struct {
	IsFinal bool
	Next    map[uint8]*Trie
}

func (t *Trie) Add(s string, i int) {
	if t == nil {
		return
	}

	if i == len(s) {
		t.IsFinal = true

		return
	}

	if t.Next == nil {
		next := &Trie{}
		t.Next = map[uint8]*Trie{
			s[i]: next,
		}

		next.Add(s, i+1)

		return
	}

	if next, ok := t.Next[s[i]]; ok {
		next.Add(s, i+1)

		return
	}

	next := &Trie{}
	t.Next[s[i]] = next

	next.Add(s, i+1)

	return
}

func (t *Trie) Present(s string, i int) []int {
	if t == nil {
		return []int{}
	}

	result := make([]int, 0)
	if t.IsFinal {
		result = append(result, i)
	}

	if t.Next == nil || i == len(s) {
		return result
	}

	next, ok := t.Next[s[i]]
	if ok {
		children := next.Present(s, i+1)
		result = append(result, children...)
	}

	return result
}

func wordBreak(s string, wordDict []string) bool {
	// we need to go though each letter and check if there is a word start nd for how long
	// after that we create "nodes" represnting letters as letters and connect nodes i and j with edges if from letter i we have a word with length k=j-i.
	// we need check if there is a path from 0 to n-1
	// time O(wlen + n*m + n*n) and memory O(n*n + wlen)

	root := &Trie{}
	n := len(s)
	for _, word := range wordDict {
		root.Add(word, 0)
	}

	nodes := make([][]bool, n+1)
	for i := 0; i < n; i++ {
		nodes[i] = make([]bool, n+1)
		edges := root.Present(s[i:], 0)
		for _, edge := range edges {
			nodes[i][i+edge] = true
		}
	}
	nodes[n] = make([]bool, n+1)

	marked := make([]bool, n+1)

	return find(nodes, marked, 0, n+1, n)
}

func find(nodes [][]bool, marked []bool, i, n, destination int) bool {
	if i == destination {
		return true
	}

	marked[i] = true
	discovered := false
	for j := 0; j < n && !discovered; j++ {
		if nodes[i][j] && !marked[j] {
			discovered = find(nodes, marked, j, n, destination)
		}
	}

	return discovered
}
