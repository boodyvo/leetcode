package task

// doesn't pass time limit in Golang but pass in C++ on HackerRank

func morganAndString(a string, b string) string {
	first, second := 0, 0
	aLen := len(a)
	bLen := len(b)
	result := ""

	for first < aLen || second < bLen {
		//fmt.Println("new check", first, second)

		if first == aLen {
			result += string(b[second])
			second++

			continue
		}

		if second == bLen {
			result += string(a[first])
			first++

			continue
		}

		if a[first] > b[second] {
			result += string(b[second])
			second++

			continue
		}

		if a[first] < b[second] {
			result += string(a[first])
			first++

			continue
		}

		if compare(a, b, first+1, second+1) {
			result += string(a[first])
			first++
			for ; first < aLen && a[first] == a[first-1]; first++ {
				result += string(a[first])
			}

			continue
		}

		result += string(b[second])
		second++
		for ; second < bLen && b[second] == b[second-1]; second++ {
			result += string(b[second])
		}
	}

	return result
}

func compare(a, b string, first, second int) bool {
	aLen := len(a)
	bLen := len(b)
	for first < aLen && second < bLen {
		if a[first] < b[second] {
			return true
		}
		if a[first] > b[second] {
			return false
		}

		first++
		second++
	}

	return first != aLen
}
