package equileader

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
			a: []int{4, 3, 4, 4, 4, 2},
			r: 2,
		},
		{
			a: []int{2},
			r: 0,
		},
		{
			a: []int{2, 4, 8, 10, 12},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}

func BenchmarkSolution100000NoLeader(b *testing.B) {
	n := 100000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	for i := 0; i < b.N; i++ {
		Solution(a)
	}
}
