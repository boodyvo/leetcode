package task

type Character struct {
	First int
	Last  int
	Char  uint8
}

func solve(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	var firstChar, secondChar *Character
	current := 0
	best := 0

	for i := 0; i < n; i++ {
		if firstChar != nil && s[i] == firstChar.Char {
			current++
			if current > best {
				best = current
			}

			continue
		}
		if secondChar != nil && s[i] == secondChar.Char {
			current++
			if current > best {
				best = current
			}

			continue
		}

		if firstChar == nil {
			firstChar = &Character{
				First: i,
				Last:  i,
				Char:  s[i],
			}

			current++
			if current > best {
				best = current
			}
		} else if secondChar == nil {
			secondChar = &Character{
				First: i,
				Last:  i,
				Char:  s[i],
			}
			current++
			if current > best {
				best = current
			}
		} else if firstChar.Last < secondChar.Last {
			secondChar.First = max(secondChar.First, firstChar.Last+1)
			firstChar = &Character{
				First: i,
				Last:  i,
				Char:  s[i],
			}
			current = i - secondChar.First + 1
			if current > best {
				best = current
			}
		} else {
			firstChar.First = max(firstChar.First, secondChar.Last+1)
			secondChar = &Character{
				First: i,
				Last:  i,
				Char:  s[i],
			}
			current = i - firstChar.First + 1
			if current > best {
				best = current
			}
		}

	}

	if current > best {
		best = current
	}

	return best
}
