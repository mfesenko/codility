package maxslicesum

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
			a: []int{3, 2, -6, 4, 0},
			r: 5,
		},
		{
			a: []int{-25},
			r: -25,
		},
		{
			a: []int{5, 6, -25, -4, 17, 0},
			r: 17,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
