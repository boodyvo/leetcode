package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		inputM         int
		inputK         int
		inputBloomDays []int
		expected       int
	}{
		//{
		//	inputM:         3,
		//	inputK:         1,
		//	inputBloomDays: []int{1, 10, 3, 10, 2},
		//	expected:       3,
		//},
		//{
		//	inputM:         3,
		//	inputK:         2,
		//	inputBloomDays: []int{1, 10, 3, 10, 2},
		//	expected:       -1,
		//},
		{
			inputM:         2,
			inputK:         3,
			inputBloomDays: []int{7, 7, 7, 7, 12, 7, 7},
			expected:       12,
		},
	}

	for _, tc := range testCases {
		actual := minDays(tc.inputBloomDays, tc.inputM, tc.inputK)
		require.Equal(t, tc.expected, actual)
	}
}
