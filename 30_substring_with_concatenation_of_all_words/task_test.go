package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		s        string
		words    []string
		expected []int
	}{
		{
			s:        "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0, 9},
		},
	}

	for _, tc := range testCases {
		actual := findSubstring(tc.s, tc.words)
		require.Equal(t, tc.expected, actual)
	}
}
