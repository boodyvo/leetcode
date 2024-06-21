package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input    []uint64
		expected uint64
	}{
		{
			input:    []uint64{2, 5, 9},
			expected: 504,
		},
	}

	for _, tc := range testCases {
		actual := findConnectedComponents(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
