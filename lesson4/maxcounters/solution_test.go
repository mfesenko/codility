package maxcounters

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	maxCounters   = 100000
	maxOperations = 100000
)

type test struct {
	n int
	a []int
	r []int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			n: 5,
			a: []int{3, 4, 4, 6, 1, 4, 4},
			r: []int{3, 2, 2, 4, 2},
		},
		{
			n: 3,
			a: []int{4, 4, 4, 4},
			r: []int{0, 0, 0},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.n, test.a))
	}
}

func BenchmarkSolution100000IncreaseOperations(b *testing.B) {
	a := make([]int, maxOperations)
	for i := 0; i < maxOperations; i++ {
		a[i] = randomIncreaseOperation(maxCounters)
	}

	runBenchmark(b, maxCounters, a)
}

func BenchmarkSolution100000MaxCounterOperations(b *testing.B) {
	a := make([]int, maxOperations)
	operation := maxCountersOperation(maxCounters)
	for i := 0; i < maxOperations; i++ {
		a[i] = operation
	}

	runBenchmark(b, maxCounters, a)
}

func BenchmarkSolution100000WithHalfMaxCounterOperations(b *testing.B) {
	a := make([]int, maxOperations)
	for i := 0; i < maxOperations; i++ {
		if i%2 == 0 {
			a[i] = randomIncreaseOperation(maxCounters)
		} else {
			a[i] = maxCountersOperation(maxCounters)
		}
	}

	runBenchmark(b, maxCounters, a)
}

func BenchmarkSolution100000RandomOperations(b *testing.B) {
	a := make([]int, maxOperations)
	for i := 0; i < maxOperations; i++ {
		a[i] = rand.Intn(maxCounters+1) + 1
	}

	runBenchmark(b, maxCounters, a)
}

func randomIncreaseOperation(n int) int {
	return rand.Intn(n) + 1
}

func maxCountersOperation(n int) int {
	return n + 1
}

func runBenchmark(b *testing.B, n int, a []int) {
	for i := 0; i < b.N; i++ {
		Solution(n, a)
	}
}
