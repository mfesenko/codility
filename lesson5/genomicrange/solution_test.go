package genomicrange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	s string
	p []int
	q []int
	r []int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			s: "CAGCCTA",
			p: []int{2, 5, 0},
			q: []int{4, 5, 6},
			r: []int{2, 4, 1},
		},
		{
			s: "CCAGCTTCAGT",
			p: []int{5, 2, 3},
			q: []int{6, 2, 6},
			r: []int{4, 1, 2},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.s, test.p, test.q))
	}
}
