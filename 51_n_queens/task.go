package task

func solveNQueens(n int) [][]string {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	results := generate(board, n, 0)
	convertedResults := make([][]string, 0)

	for _, result := range results {
		convertedResults = append(convertedResults, convert(result))
	}

	return convertedResults
}

func generate(board [][]int, n, line int) [][][]int {
	// final
	if line == n {
		return [][][]int{board}
	}

	result := make([][][]int, 0)

	for i := 0; i < n; i++ {
		if board[line][i] == 0 {
			var newBoard [][]int
			for _, row := range board {
				newRow := make([]int, len(row))
				copy(newRow, row)
				newBoard = append(newBoard, newRow)
			}

			newBoard[line][i] = 1
			for j := 0; j < n; j++ {
				if newBoard[line][j] == 0 {
					newBoard[line][j] = -1
				}
				if newBoard[j][i] == 0 {
					newBoard[j][i] = -1
				}
				if line+j < n && i+j < n && newBoard[line+j][i+j] == 0 {
					newBoard[line+j][i+j] = -1
				}
				if line-j >= 0 && i-j >= 0 && newBoard[line-j][i-j] == 0 {
					newBoard[line-j][i-j] = -1
				}
				if line+j < n && i-j >= 0 && newBoard[line+j][i-j] == 0 {
					newBoard[line+j][i-j] = -1
				}
				if line-j >= 0 && i+j < n && newBoard[line-j][i+j] == 0 {
					newBoard[line-j][i+j] = -1
				}
			}

			solutions := generate(newBoard, n, line+1)

			result = append(result, solutions...)
		}
	}

	return result
}

func convert(board [][]int) []string {
	result := make([]string, 0)

	for _, row := range board {
		rowStr := ""
		for _, cell := range row {
			if cell == 1 {
				rowStr += "Q"
			} else {
				rowStr += "."
			}
		}
		result = append(result, rowStr)
	}

	return result
}
