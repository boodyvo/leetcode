package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{1, 1, 0}, {1, 1, 1}, {1, 2, 1}},
			expected: 3,
		},
		{
			input:    [][]int{{1, 3, 0}, {1, 0, 0}, {1, 0, 3}},
			expected: 4,
		},
		{
			input:    [][]int{{0, 0, 0}, {0, 9, 0}, {0, 0, 0}},
			expected: 12,
		},
	}

	for _, tc := range testCases {
		actual := minimumMoves(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
