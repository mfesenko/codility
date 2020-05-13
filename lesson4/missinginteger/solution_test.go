package missinginteger

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
			a: []int{1, 3, 6, 4, 1, 2},
			r: 5,
		},
		{
			a: []int{1, 2, 3},
			r: 4,
		},
		{
			a: []int{-1, -3},
			r: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
