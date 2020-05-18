package counttriangles

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
			a: []int{10, 2, 5, 1, 8, 12},
			r: 4,
		},
		{
			a: nil,
			r: 0,
		},
		{
			a: []int{3, 4},
			r: 0,
		},
		{
			a: []int{5, 3, 4},
			r: 1,
		},
		{
			a: []int{13, 4, 5},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
