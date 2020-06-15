package flooddepth

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
			a: []int{1, 3, 2, 1, 2, 1, 5, 3, 3, 4},
			r: 2,
		},
		{
			a: []int{5, 8},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
