package passingcars

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
			a: []int{0, 1, 0, 1, 1},
			r: 5,
		},
		{
			a: []int{1, 1, 0, 0, 0},
			r: 0,
		},
		{
			a: []int{0, 0, 0, 1, 1},
			r: 6,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}

func TestSolutionReturnsMinusOneWhenTheAmountOfPassingCarsIsTooHigh(t *testing.T){
	n := 100000
	a := make([]int, n)
	for i := 20000; i < n; i++ {
		a[i] = 1
	}
	assert.Equal(t, tooManyPassingCars, Solution(a))
}
