package tapeequilibrium

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
			a: []int{3, 1, 2, 4, 3},
			r: 1,
		},
		{
			a: []int{-1000, 1000},
			r: 2000,
		},
		{
			a: []int{3, -4, 2, -2},
			r: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
