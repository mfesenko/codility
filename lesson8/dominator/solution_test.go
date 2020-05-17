package dominator

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
			a: []int{3, 4, 3, 2, 3, -1, 3, 3},
			r: 7,
		},
		{
			a: nil,
			r: -1,
		},
		{
			a: []int{5, 2, 5, 2, 5, 2},
			r: -1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
