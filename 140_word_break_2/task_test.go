package task

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		input     string
		inputDict []string
		expected  []string
	}{
		{
			input:     "catsanddog",
			inputDict: []string{"cat", "cats", "and", "sand", "dog"},
			expected:  []string{"cats and dog", "cat sand dog"},
		},
		{
			input:     "catsandog",
			inputDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected:  []string{},
		},
	}

	for _, tc := range testCases {
		actual := wordBreak(tc.input, tc.inputDict)
		sort.Strings(tc.expected)
		sort.Strings(actual)
		require.ElementsMatch(t, tc.expected, actual)
	}
}
