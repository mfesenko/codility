package maxnonoverlappingsegments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a []int
	b []int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: []int{1, 3, 7, 9, 9},
			b: []int{5, 6, 8, 9, 10},
			r: 3,
		},
		{
			a: nil,
			b: nil,
			r: 0,
		},
		{
			a: []int{1, 1, 0, 2},
			b: []int{5, 4, 2, 3},
			r: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a, test.b))
	}
}
