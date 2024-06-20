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
			input:    "{[()]}",
			expected: "YES",
		},
		{
			input:    "{[({)}]}",
			expected: "NO",
		},
	}

	for _, tc := range testCases {
		actual := solve(tc.input)
		require.Equal(t, tc.expected, actual)
	}
}
