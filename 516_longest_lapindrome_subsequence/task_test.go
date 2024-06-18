package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "babad",
			expected: 3,
		},
		{
			input:    "cbbd",
			expected: 2,
		},
	}

	for _, tc := range testCases {
		actual := longestPalindromeSubseq(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
