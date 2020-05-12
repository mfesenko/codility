package oddoccurences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a []int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: []int{9, 3, 9, 3, 9, 7, 9},
			r: 7,
		},
		{
			a: []int{1, 1, 1},
			r: 1,
		},
		{
			a: []int{1, 1},
			r: -1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
