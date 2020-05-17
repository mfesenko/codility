package minperimeter

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
			n: 30,
			r: 22,
		},
		{
			n: 1,
			r: 4,
		},
		{
			n: 16,
			r: 16,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.n))
	}
}
