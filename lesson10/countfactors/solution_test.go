package countfactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	n int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			n: 24,
			r: 8, // 1, 2, 3, 4, 6, 8, 12, 24
		},
		{
			n: 16,
			r: 5, // 1, 2, 4, 8, 16
		},
		{
			n: 1,
			r: 1, // 1
		},
		{
			n: 17,
			r: 2, // 1, 17
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.n))
	}
}
