package countsemiprimes

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	n int
	p []int
	q []int
	r []int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			n: 26,
			p: []int{1, 4, 16},
			q: []int{26, 10, 20},
			r: []int{10, 4, 0},
		},
		{
			n: 50,
			p: []int{17, 34},
			q: []int{17, 34},
			r: []int{0, 1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.n, test.p, test.q))
	}
}

func BenchmarkSolution(b *testing.B) {
	n := 50000
	m := 30000
	p, q := generateRandomInput(n, m)
	for i := 0; i < b.N; i++ {
		Solution(n, p, q)
	}
}

func generateRandomInput(n int, m int) ([]int, []int) {
	p := make([]int, m)
	q := make([]int, m)
	for i := 0; i < m; i++ {
		start := rand.Intn(n)
		end := rand.Intn(n)
		if end < start {
			start, end = end, start
		}
		p[i] = start
		q[i] = end
	}
	return p, q
}
