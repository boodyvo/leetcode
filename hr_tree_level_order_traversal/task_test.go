package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		inputN     int
		inputNodes []int
		expected   []int
	}{
		{
			inputN:     6,
			inputNodes: []int{1, 2, 5, 3, 4, 6},
			expected:   []int{1, 2, 5, 3, 6, 4},
		},
	}

	for _, tc := range testCases {
		actual := solve(tc.inputN, tc.inputNodes)
		require.Equal(t, tc.expected, actual)
	}
}
