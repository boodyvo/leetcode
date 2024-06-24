package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		word      string
		forbidden []string
		expected  int
	}{
		{
			word:      "cbaaaabc",
			forbidden: []string{"aaa", "cb"},
			expected:  4,
		},
		{
			word:      "leetcode",
			forbidden: []string{"de", "le", "e"},
			expected:  4,
		},
	}

	for _, tc := range testCases {
		actual := longestValidSubstring(tc.word, tc.forbidden)
		require.Equal(t, tc.expected, actual)
	}
}
