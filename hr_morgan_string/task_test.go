package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		inputA   string
		inputB   string
		expected string
	}{
		{
			inputA:   "ABC",
			inputB:   "ABCAAC",
			expected: "AABBCAACC",
		},
		{
			inputA:   "AADDBBCC",
			inputB:   "AADDBBCCDDAA",
			expected: "AAAADDBBCCDDAADDBBCC",
		},
		{
			inputA:   "JACK",
			inputB:   "DANIEL",
			expected: "DAJACKNIEL",
		},
		{
			inputA:   "ABACABA",
			inputB:   "ABACABA",
			expected: "AABABACABACABA",
		},
	}

	for _, tc := range testCases {
		actual := morganAndString(tc.inputA, tc.inputB)
		require.Equal(t, tc.expected, actual)
	}
}
