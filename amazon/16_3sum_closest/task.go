package task

import "sort"

const INF = 100_000_000

func abs(a, b int) int {
	if a >= b {
		return a - b
	}

	return b - a
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	numbers := make(map[int]int, n)
	for i := 0; i < n; i++ {
		numbers[nums[i]]++
	}

	best := INF
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			toFind := target - nums[i] - nums[j]
			pos := find(toFind, nums)
			for k := max(pos-2, 0); k < pos+2 && k < n; k++ {
				needed := nums[k]
				amount := 1
				if nums[i] == needed {
					amount++
				}
				if nums[j] == needed {
					amount++
				}
				if numbers[needed] < amount {
					continue
				}

				if abs(target, nums[i]+nums[j]+nums[k]) < abs(target, best) {
					best = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}

	return best
}

func find(target int, numbers []int) int {
	left := 0
	right := len(numbers)
	for left < right {
		mid := (left + right) / 2
		if numbers[mid] == target {
			return mid
		}

		if numbers[mid] > target {
			right = mid

			continue
		}

		left = mid + 1
	}

	return left
}
