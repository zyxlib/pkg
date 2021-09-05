package iter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermutation(t *testing.T) {
	testCases := []struct {
		name  string
		testF func(*testing.T)
	}{
		{
			name: "Zero",
			testF: func(t *testing.T) {
				perm := PermutationInts(nil)
				require.Len(t, perm, 0)
			},
		},
		{
			name: "One",
			testF: func(t *testing.T) {
				perm := PermutationInts([]int{0})
				require.Len(t, perm, 1)
				require.Contains(t, perm, []int{0})
			},
		},
		{
			name: "Two",
			testF: func(t *testing.T) {
				perm := PermutationInts([]int{0, 1})
				require.Len(t, perm, 2)
				require.Contains(t, perm, []int{0, 1})
				require.Contains(t, perm, []int{1, 0})
			},
		},
		{
			name: "Three",
			testF: func(t *testing.T) {
				perm := PermutationInts([]int{0, 1, 2})
				require.Len(t, perm, 6)
				require.Contains(t, perm, []int{0, 1, 2})
				require.Contains(t, perm, []int{0, 2, 1})
				require.Contains(t, perm, []int{1, 0, 2})
				require.Contains(t, perm, []int{1, 2, 0})
				require.Contains(t, perm, []int{2, 0, 1})
				require.Contains(t, perm, []int{2, 1, 0})
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, tc.testF)
	}
}
