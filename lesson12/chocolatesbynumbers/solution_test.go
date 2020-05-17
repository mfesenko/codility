package chocolatesbynumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	n int
	m int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			n: 10,
			m: 4,
			r: 5,
		},
		{
			n: 1,
			m: 5,
			r: 1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.n, test.m))
	}
}
