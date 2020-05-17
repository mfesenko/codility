package commonprimedivisors

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
			a: []int{15, 10, 3},
			b: []int{75, 30, 5},
			r: 1,
		},
		{
			a: []int{15, 4, 16},
			b: []int{45, 30, 64},
			r: 2,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a, test.b))
	}
}
