package task

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	testCases := []struct {
		n, s        int
		genes       []string
		healths     []int
		strands     []Strand
		expectedMin int
		expectedMax int
	}{
		{
			n:       6,
			s:       3,
			genes:   []string{"a", "b", "c", "aa", "d", "b"},
			healths: []int{1, 2, 3, 4, 5, 6},
			strands: []Strand{
				{Left: 1, Right: 5, Value: "caaab"},
				{Left: 0, Right: 4, Value: "xyz"},
				{Left: 2, Right: 4, Value: "bcdybc"},
			},
			expectedMin: 0,
			expectedMax: 19,
		},
		{
			n:       6,
			s:       3,
			genes:   []string{"a", "aa", "aaa", "aaaa", "aaaaa", "a"},
			healths: []int{1, 2, 3, 4, 5, 6},
			strands: []Strand{
				{Left: 1, Right: 5, Value: "aaaaaaaaaa"},
				{Left: 0, Right: 5, Value: "a"},
				{Left: 1, Right: 5, Value: "a"},
				{Left: 0, Right: 4, Value: "a"},
			},
			expectedMin: 1,
			expectedMax: 160,
		},
		{
			n:       6,
			s:       3,
			genes:   []string{"a", "aa", "aaa", "aaaa", "aaaaa", "a"},
			healths: []int{1, 2, 3, 4, 5, 6},
			strands: []Strand{
				{Left: 1, Right: 5, Value: "aaaaaaaaaa"},
			},
			expectedMin: 160,
			expectedMax: 160,
		},
	}

	for _, tc := range testCases {
		resMin, resMax := solve(tc.n, tc.s, tc.genes, tc.healths, tc.strands)
		require.Equal(t, tc.expectedMin, resMin)
		require.Equal(t, tc.expectedMax, resMax)
	}
}
