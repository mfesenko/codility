package countnondivisible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a []int
	r []int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: []int{3, 1, 2, 3, 6},
			r: []int{2, 4, 3, 2, 0},
		},
		{
			a: []int{1, 2, 3, 4, 5, 6, 7},
			r: []int{6, 5, 5, 4, 5, 3, 5},
		},
		{
			a: []int{2, 1},
			r: []int{0, 1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}

func Benchmark(b *testing.B) {
	n := 50000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = (i + 1) * 2
	}
	for i := 0; i < b.N; i++ {
		Solution(a)
	}
}
