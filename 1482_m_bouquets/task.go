package task

import (
	"fmt"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	sortedDays := make([]int, len(bloomDay))
	copy(sortedDays, bloomDay)

	sort.Ints(sortedDays)

	days, found := find(m, k, 0, len(bloomDay)-1, bloomDay, sortedDays)
	if !found {
		return -1
	}

	return days
}

func find(m, k, left, right int, bloomDays, bloomDaysSorted []int) (int, bool) {
	best := -1
	for left <= right {
		cur := (left + right) / 2
		day := bloomDaysSorted[cur]

		bouquets := check(k, day, bloomDays)

		fmt.Println("day check", day, bouquets)

		if bouquets >= m {
			best = day

			right = cur - 1
		} else {
			left = cur + 1
		}
	}

	if best == -1 {
		return 0, false
	}

	return best, true
}

func check(k, day int, bloomDays []int) int {
	counter := 0
	cur := 0
	for _, bloom := range bloomDays {
		if bloom <= day {
			cur++
		} else {
			cur = 0
		}

		fmt.Println("current checked", cur, counter)

		if cur/k > 0 {
			counter++
			cur = 0
		}
	}

	return counter
}
