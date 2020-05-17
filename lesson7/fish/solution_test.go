package fish

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a []int
	b []int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: []int{20},
			b: []int{1},
			r: 1,
		},
		{
			a: []int{4, 3, 2, 1, 5},
			b: []int{0, 1, 0, 0, 0},
			r: 2,
		},
		{
			a: []int{4, 3, 2, 1, 5},
			b: []int{0, 0, 0, 1, 1},
			r: 5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a, test.b))
	}
}
