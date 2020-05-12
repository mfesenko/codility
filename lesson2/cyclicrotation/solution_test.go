package cyclicrotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a []int
	k int
	r []int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: []int{3, 8, 9, 7, 6},
			k: 3,
			r: []int{9, 7, 6, 3, 8},
		},
		{
			a: []int{0, 0, 0},
			k: 1,
			r: []int{0, 0, 0},
		},
		{
			a: []int{1, 2, 3, 4},
			k: 4,
			r: []int{1, 2, 3, 4},
		},
		{
			a: []int{1, 2, 3, 4},
			k: 13,
			r: []int{4, 1, 2, 3},
		},
		{
			a: nil,
			k: 2,
			r: nil,
		},
		{
			a: []int{1},
			k: 15,
			r: []int{1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a, test.k))
	}
}
