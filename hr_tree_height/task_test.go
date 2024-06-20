package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		inputN     int
		inputNodes []int
		expected   int
	}{
		{
			inputN:     7,
			inputNodes: []int{3, 2, 1, 5, 4, 6, 7},
			expected:   3,
		},
	}

	for _, tc := range testCases {
		actual := solve(tc.inputN, tc.inputNodes)
		require.Equal(t, tc.expected, actual)
	}
}
