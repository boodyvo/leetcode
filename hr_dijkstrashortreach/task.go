package task

const INF int32 = 2_000_000_000

func shortestReach(n int32, edges [][]int32, s int32) []int32 {
	distances := make([]int32, n)
	matrix := make([][]int32, n*n)
	for i := 0; i < int(n); i++ {
		distances[i] = INF
		matrix[i] = make([]int32, n)
		for j := 0; j < int(n); j++ {
			matrix[i][j] = INF
		}
	}

	for i := 0; i < len(edges); i++ {
		if matrix[edges[i][0]-1][edges[i][1]-1] > edges[i][2] {
			matrix[edges[i][0]-1][edges[i][1]-1] = edges[i][2]
			matrix[edges[i][1]-1][edges[i][0]-1] = edges[i][2]
		}
	}

	index := int(s - 1)
	distances[index] = 0
	nodesToVisit := make([]int, 0, n)
	nodesToVisit = append(nodesToVisit, index)
	visited := make(map[int]struct{}, n)

	for len(visited) < len(nodesToVisit) {
		index = -1
		for _, node := range nodesToVisit {
			if _, ok := visited[node]; ok {
				continue
			}

			if index == -1 || distances[node] < distances[index] {
				index = node
			}
		}
		visited[index] = struct{}{}

		for i := 0; i < int(n); i++ {
			if distances[i] > distances[index]+matrix[index][i] {
				// we met this first time, so need to add to nodes for visiting
				if distances[i] == INF {
					nodesToVisit = append(nodesToVisit, i)
				}

				distances[i] = distances[index] + matrix[index][i]
			}
		}
	}

	results := make([]int32, 0, n-1)
	for i := 0; i < int(n); i++ {
		if i == int(s-1) {
			continue
		}

		if distances[i] == INF {
			results = append(results, -1)
		} else {
			results = append(results, distances[i])
		}
	}

	return results
}
