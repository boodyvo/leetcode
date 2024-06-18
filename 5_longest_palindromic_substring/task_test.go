package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
	}

	for _, tc := range testCases {
		actual := longestPalindrome(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
