package task

var (
	longestData [][]int
)

func longestPalindromeSubseq(s string) int {
	n := len(s)
	longestData = make([][]int, n)

	for i := 0; i < n; i++ {
		longestData[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j < i {
				longestData[i][j] = 0
				continue
			} else if j == i {
				longestData[i][j] = 1
				continue

			}

			longestData[i][j] = -1
		}
	}

	return solve(0, len(s)-1, s)

	// return recoverSolution(0, len(s)-1, s)
}

func solve(i, j int, s string) int {
	if longestData[i][j] != -1 {
		return longestData[i][j]
	}

	if s[i] == s[j] {
		longestData[i][j] = solve(i+1, j-1, s) + 2

		return longestData[i][j]
	}

	longestData[i][j] = max(solve(i+1, j, s), solve(i, j-1, s))

	return longestData[i][j]
}

func recoverSolution(i, j int, s string) string {
	if i > j {
		return ""
	}

	if i == j {
		return string(s[i])
	}

	if s[i] == s[j] {
		return string(s[i]) + recoverSolution(i+1, j-1, s) + string(s[j])
	}

	if longestData[i][j] == longestData[i+1][j] {
		return recoverSolution(i+1, j, s)
	}

	return recoverSolution(i, j-1, s)
}
