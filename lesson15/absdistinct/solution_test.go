package absdistinct

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
			a: []int{-5, -3, -1, 0, 3, 6},
			r: 5,
		},
		{
			a: []int{3},
			r: 1,
		},
		{
			a: []int{-3, -3, 3},
			r: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
