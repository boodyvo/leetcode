package task

type Stack struct {
	Values []int
}

func (s *Stack) Pop() {
	s.Values = s.Values[:len(s.Values)-1]
}

func (s *Stack) Top() (int, bool) {
	if len(s.Values) == 0 {
		return 0, true
	}

	return s.Values[len(s.Values)-1], false
}

func (s *Stack) Add(value int) {
	s.Values = append(s.Values, value)
}

func solve(s string) string {
	stack := Stack{}
	valueMap := map[string]int{
		"{": 1,
		"}": -1,
		"[": 2,
		"]": -2,
		"(": 3,
		")": -3,
	}

	for i := 0; i < len(s); i++ {
		val := valueMap[string(s[i])]
		if val > 0 {
			stack.Add(val)

			continue
		}

		top, empty := stack.Top()
		if empty {
			return "NO"
		}

		if top+val != 0 {
			return "NO"
		}

		stack.Pop()
	}

	if _, empty := stack.Top(); empty {
		return "YES"
	}

	return "NO"
}
