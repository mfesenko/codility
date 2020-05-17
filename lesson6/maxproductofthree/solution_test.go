package maxproductofthree

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
			a: []int{-3, 1, 2, -2, 5, 6},
			r: 60,
		},
		{
			a: []int{2, -5, 1},
			r: -10,
		},
		{
			a: []int{-2, -5, -1},
			r: -10,
		},
		{
			a: []int{-2, -5, -1, 0},
			r: 0,
		},
		{
			a: []int{10, 20, 30, 1, 2, 3, -5},
			r: 6000,
		},
		{
			a: []int{10, 20, 30, 1, 2, 3, -50, -100},
			r: 150000,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
