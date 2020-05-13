package frogriver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a []int
	x int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: []int{1, 3, 1, 4, 2, 3, 5, 4},
			x: 5,
			r: 6,
		},
		{
			a: []int{1, 3, 1, 4, 2, 3, 5, 4},
			x: 10,
			r: -1,
		},
		{
			a: []int{1, 2, 3, 4, 5, 6},
			x: 6,
			r: 5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.x, test.a))
	}
}
