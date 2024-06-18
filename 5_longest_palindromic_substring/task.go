package task

var (
	isPalindrome [][]int
)

func longestPalindrome(s string) string {
	n := len(s)
	isPalindrome = make([][]int, n)
	bestLeft := 0
	bestRight := 0

	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j < i {
				isPalindrome[i][j] = 1
				continue
			} else if j == i {
				isPalindrome[i][j] = 1
				continue
			}

			isPalindrome[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if solve(i, j, s) {
				if j-i > bestRight-bestLeft {
					bestLeft = i
					bestRight = j
				}
			}
		}
	}

	return s[bestLeft : bestRight+1]
}

func solve(i, j int, s string) bool {
	if isPalindrome[i][j] != -1 {
		return isPalindrome[i][j] == 1
	}

	if s[i] == s[j] {
		result := solve(i+1, j-1, s)
		if result {
			isPalindrome[i][j] = 1
			return true
		}

		isPalindrome[i][j] = 0

		return false
	}

	isPalindrome[i][j] = 0

	return false
}
