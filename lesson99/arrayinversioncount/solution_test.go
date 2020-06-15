package arrayinversioncount

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
			a: []int{-1, 6, 3, 4, 7, 4},
			r: 4,
		},
		{
			a: []int{1, 2, 3, 4},
			r: 0,
		},
		{
			a: nil,
			r: 0,
		},
		{
			a: []int{2, 1, 3, 4},
			r: 1,
		},
		{
			a: []int{5, 5, 5},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}

func TestSolutionReturnsMinusOneWhenInversionCountIsHigherThanMaxValue(t *testing.T) {
	n := 100000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = n - i
	}
	assert.Equal(t, -1, Solution(a))
}
