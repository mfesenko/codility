package minavgtwoslice

import (
	"math/rand"
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
			a: []int{4, 2, 2, 5, 1, 5, 8},
			r: 1,
		},
		{
			a: []int{10, 12, 6, 5, 4},
			r: 3,
		},
		{
			a: []int{7, 6, 5, 4, 10, 12, -32, -2, -20},
			r: 6,
		},
		{
			a: []int{7, 6, 5, 4, 10, 7, 2},
			r: 2,
		},
		{
			a: []int{-10, 3},
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a))
	}
}

func Benchmark100000(b *testing.B) {
	n := 100000
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(20000) - 10000
	}
	for i := 0; i < b.N; i++ {
		Solution(a)
	}
}
