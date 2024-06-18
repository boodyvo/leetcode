package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input    int
		expected [][]string
	}{
		{
			input:    4,
			expected: [][]string{{"..Q.", "Q...", "...Q", ".Q.."}, {".Q..", "...Q", "Q...", "..Q."}},
		},
		{
			input:    1,
			expected: [][]string{{"Q"}},
		},
	}

	for _, tc := range testCases {
		actual := solveNQueens(tc.input)
		require.Equal(t, len(tc.expected), len(actual))
	}
}
