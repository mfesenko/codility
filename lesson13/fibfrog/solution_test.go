package fibfrog

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
			a: []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0},
			r: 3,
		},
		{
			a: []int{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
			r: -1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}
