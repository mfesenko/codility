package tieropes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	k int
	a []int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			k: 4,
			a: []int{1, 2, 3, 4, 1, 1, 3},
			r: 3,
		},
		{
			k: 20,
			a: []int{1, 2, 3, 4},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.k, test.a))
	}
}
