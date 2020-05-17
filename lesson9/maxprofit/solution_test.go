package maxprofit

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
			a: []int{23171, 21011, 21123, 21366, 21013, 21367},
			r: 356,
		},
		{
			a: nil,
			r: 0,
		},
		{
			a: []int{25678, 25677, 23500, 12000},
			r: 0,
		},
		{
			a: []int{23, 24, 19, 10, 15},
			r: 5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
