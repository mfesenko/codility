package permcheck

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
			a: []int{4, 1, 3, 2},
			r: 1,
		},
		{
			a: []int{4, 1, 3},
			r: 0,
		},
		{
			a: []int{50000, 1, 3, 2},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
