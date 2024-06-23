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
			input:    "ccaabbb",
			expected: 5,
		},
		{
			input:    "ccaabb",
			expected: 4,
		},
		{
			input:    "aaaaaaaa",
			expected: 8,
		},
		{
			input:    "abcabcabcabc",
			expected: 2,
		},
		{
			input:    "ab",
			expected: 2,
		},
		{
			input:    "a",
			expected: 1,
		},
		{
			input:    "",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := solve(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
