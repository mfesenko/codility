package distinct

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
			a: []int{2, 1, 1, 2, 3, 1},
			r: 3,
		},
		{
			a: nil,
			r: 0,
		},
		{
			a: []int{5, 5, 5, 5},
			r: 1,
		},
		{
			a: []int{-7, 5, 1, 3},
			r: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
