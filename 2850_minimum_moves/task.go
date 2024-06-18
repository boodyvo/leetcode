package task

import (
	"math"
)

type Coordinate struct {
	X, Y int
}

func minimumMoves(grid [][]int) int {
	destination := make([]Coordinate, 0, 9)
	sources := make([]Coordinate, 0, 9)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				destination = append(destination, Coordinate{i, j})
			} else if grid[i][j] > 1 {
				for count := 1; count < grid[i][j]; count++ {
					sources = append(sources, Coordinate{i, j})
				}
			}
		}
	}

	return solve(destination, sources)
}

func solve(destination, source []Coordinate) int {
	used := make(map[int]bool, len(destination))
	shuffle := make([]int, len(destination))

	return findBest(0, used, shuffle, destination, source)
}

func findBest(index int, used map[int]bool, shuffle []int, destination, source []Coordinate) int {
	if index == len(destination) {
		return distanceTotal(shuffle, destination, source)
	}

	best := 10000
	for i := 0; i < len(destination); i++ {
		if used[i] {
			continue
		}

		used[i] = true
		shuffle[index] = i
		potentialBest := findBest(index+1, used, shuffle, destination, source)
		if potentialBest < best {
			best = potentialBest
		}
		used[i] = false
	}

	return best
}

func distanceTotal(shuffle []int, destination, source []Coordinate) int {
	sum := 0
	for i := 0; i < len(destination); i++ {
		sum += distance(destination[i], source[shuffle[i]])
	}

	return sum
}

func distance(destination, source Coordinate) int {
	return int(math.Abs(float64(destination.Y-source.Y))) + int(math.Abs(float64(destination.X-source.X)))
}
