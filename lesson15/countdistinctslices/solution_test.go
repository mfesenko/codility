package countdistinctslices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	m int
	a []int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			m: 6,
			a: []int{3, 4, 5, 5, 2},
			r: 9,
		},
		{
			m: 2,
			a: []int{1, 1, 1, 1},
			r: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.m, test.a))
	}
}

func TestSolutionReturnsMaxValueWhenSliceCountIsHigherThanMaxValue(t *testing.T) {
	n := 100000
	m := 100000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	assert.Equal(t, maxSliceCount, Solution(m, a))
}
