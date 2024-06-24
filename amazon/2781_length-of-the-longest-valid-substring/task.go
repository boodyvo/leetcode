package task

const maxDepth = 10

type Trie struct {
	Children map[uint8]*Trie
	Final    bool
}

// AddNode adds node in reverse order
func (t *Trie) AddNode(s string, index int) {
	if t == nil {
		return
	}

	if index == len(s) {
		t.Final = true

		return
	}

	pos := len(s) - index - 1

	if t.Children == nil {
		child := &Trie{}
		t.Children = map[uint8]*Trie{
			s[pos]: child,
		}

		child.AddNode(s, index+1)

		return
	}

	if child, ok := t.Children[s[pos]]; ok {
		child.AddNode(s, index+1)

		return
	}

	child := &Trie{}
	t.Children[s[pos]] = child

	child.AddNode(s, index+1)
}

// Present checks presence in reverse order
func (t *Trie) Present(s string, index int) (int, bool) {
	if t.Final {
		return index, true
	}

	if index == len(s) {
		return index, t.Final
	}

	if t == nil || t.Children == nil {
		return 0, false
	}

	pos := len(s) - index - 1

	if child, ok := t.Children[s[pos]]; ok {
		return child.Present(s, index+1)
	}

	return 0, false
}

func longestValidSubstring(word string, forbidden []string) int {
	n := len(word)
	root := &Trie{}
	for _, forbiddenWord := range forbidden {
		root.AddNode(forbiddenWord, 0)
	}

	left := 0
	best := 0

	for right := 1; right <= n; right++ {
		if right < left {
			continue
		}

		start := left
		if right-maxDepth > start {
			start = right - maxDepth
		}

		if startedForbidden, found := root.Present(word[start:right], 0); found {
			left = right - startedForbidden + 1
		}

		best = max(best, right-left)
	}

	return best
}
