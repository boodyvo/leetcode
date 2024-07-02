package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
		{
			nums:     []int{0, 0, 0},
			target:   1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := threeSumClosest(tc.nums, tc.target)
		require.Equal(t, tc.expected, actual)
	}
}
