package numbersolitaire

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
			a: []int{1, -2, 0, 9, -1, -2},
			r: 8,
		},
		{
			a: []int{-2, 5, 1},
			r: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
