package task

import (
	"math/bits"
)

func findConnectedComponents(d []uint64) uint64 {
	used := make([]bool, len(d))

	return calculateGraph(used, d, 0)
}

func calculateGraph(used []bool, d []uint64, index int) uint64 {
	if index == len(d) {
		res := calculateComponents(used, d)

		return res
	}

	used[index] = true
	with := calculateGraph(used, d, index+1)
	used[index] = false
	without := calculateGraph(used, d, index+1)

	return with + without
}

func calculateComponents(used []bool, d []uint64) uint64 {
	checked := make([]bool, len(used))
	for i := 0; i < len(used); i++ {
		checked[i] = used[i]
	}
	data := make([]uint64, len(d))
	for i := 0; i < len(d); i++ {
		if checked[i] {
			data[i] = d[i]
		}
	}

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]&data[j] > 0 {
				data[i] = data[i] | data[j]
				data[j] = 0

				checked[j] = false
			}
		}
	}

	separateComponents := uint64(0)
	totalBits := uint64(0)
	for i := 0; i < len(data); i++ {
		if data[i] > 0 {
			totalBits += uint64(bits.OnesCount64(data[i]))
			separateComponents++
		}
	}

	return 64 - totalBits + separateComponents
}
