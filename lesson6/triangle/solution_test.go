package triangle

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
			a: []int{10, 2, 5, 1, 8, 20},
			r: 1,
		},
		{
			a: []int{10, 50, 5, 1},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
