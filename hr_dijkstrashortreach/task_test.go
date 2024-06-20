package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		inputN     int32
		inputM     int32
		inputS     int32
		inputEdges [][]int32
		expected   []int32
	}{
		{
			inputN: 4,
			inputM: 4,
			inputS: 1,
			inputEdges: [][]int32{
				{1, 2, 24},
				{1, 4, 20},
				{3, 1, 3},
				{4, 3, 12},
			},
			expected: []int32{24, 3, 15},
		},
		{
			inputN:     4,
			inputM:     0,
			inputS:     1,
			inputEdges: [][]int32{},
			expected:   []int32{-1, -1, -1},
		},
		{
			inputN: 4,
			inputM: 3,
			inputS: 1,
			inputEdges: [][]int32{
				{1, 2, 1},
				{1, 3, 1},
				{1, 4, 1},
			},
			expected: []int32{1, 1, 1},
		},
		{
			inputN: 4,
			inputM: 3,
			inputS: 4,
			inputEdges: [][]int32{
				{1, 2, 1},
				{1, 3, 1},
				{1, 4, 1},
			},
			expected: []int32{1, 2, 2},
		},
		{
			inputN: 20,
			inputM: 54,
			inputS: 17,
			inputEdges: [][]int32{
				{1, 7, 45},
				{2, 14, 15},
				{3, 7, 29},
				{4, 1, 48},
				{5, 1, 66},
				{6, 7, 17},
				{7, 14, 15},
				{8, 14, 43},
				{9, 1, 27},
				{10, 1, 33},
				{11, 14, 64},
				{12, 14, 27},
				{13, 7, 66},
				{14, 7, 54},
				{15, 14, 56},
				{16, 7, 21},
				{17, 1, 20},
				{18, 1, 34},
				{19, 7, 52},
				{20, 14, 14},
				{9, 14, 9},
				{15, 1, 39},
				{12, 1, 24},
				{9, 1, 16},
				{1, 2, 33},
				{18, 1, 46},
				{9, 1, 28},
				{15, 14, 3},
				{12, 1, 27},
				{1, 2, 5},
				{15, 1, 34},
				{1, 2, 28},
				{9, 7, 16},
				{3, 7, 23},
				{9, 7, 21},
				{9, 14, 19},
				{3, 1, 20},
				{3, 1, 5},
				{12, 14, 19},
				{3, 14, 2},
				{12, 1, 46},
				{3, 14, 5},
				{9, 14, 44},
				{6, 14, 26},
				{9, 14, 16},
				{9, 14, 34},
				{6, 7, 42},
				{3, 14, 27},
				{1, 7, 9},
				{1, 7, 41},
				{15, 14, 19},
				{12, 7, 13},
				{3, 7, 10},
				{1, 7, 2},
			},
			expected: []int32{20, 25, 25, 68, 86, 39, 22, 70, 36, 53, 91, 35, 88, 27, 30, 43, 54, 74, 41},
		},
	}

	for _, tc := range testCases {
		actual := shortestReach(tc.inputN, tc.inputEdges, tc.inputS)
		require.Equal(t, tc.expected, actual)
	}
}
