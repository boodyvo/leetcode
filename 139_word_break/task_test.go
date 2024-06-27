package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input     string
		inputDict []string
		expected  bool
	}{
		{
			input:     "leetcode",
			inputDict: []string{"leet", "code"},
			expected:  true,
		},
		{
			input:     "catsandog",
			inputDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected:  false,
		},
	}

	for _, tc := range testCases {
		actual := wordBreak(tc.input, tc.inputDict)
		require.Equal(t, tc.expected, actual)
	}
}
