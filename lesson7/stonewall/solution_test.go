package stonewall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	h []int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			h: []int{8, 8, 5, 7, 9, 8, 7, 4, 8},
			r: 7,
		},
		{
			h: []int{9},
			r: 1,
		},
		{
			h: []int{7, 7, 7},
			r: 1,
		},
		{
			h: []int{5, 5, 5, 6, 7, 6, 5, 5},
			r: 3,
		},
		{
			h: []int{5, 5, 5, 6, 7, 6, 3, 3},
			r: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.h))
	}
}
