package discintersections

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
			a: []int{1, 5, 2, 1, 4, 0},
			r: 11,
		},
		{
			a: nil,
			r: 0,
		},
		{
			a: []int{0, 0, 0, 0},
			r: 0,
		},
		{
			a: []int{1, 0, 0, 1, 0, 0},
			r: 3,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}

func TestSolutionReturnsMinusOneWhenTheAmountOfIntersectionsIsTooHigh(t *testing.T) {
	n := 100000
	a := make([]int, n)
	for i := 0; i < n; i++{
		a[i] = 10000 + i
	}
	assert.Equal(t, -1, Solution(a))
}
