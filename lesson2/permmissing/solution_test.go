package permmissing

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
			a: []int{2, 3, 1, 5},
			r: 4,
		},
		{
			a: nil,
			r: 1,
		},
		{
			a: []int{4, 3, 2, 1},
			r: 5,
		},
		{
			a: []int{4, 3, 2, 5},
			r: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
