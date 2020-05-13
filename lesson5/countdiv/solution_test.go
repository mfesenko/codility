package countdiv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	a int
	b int
	k int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			a: 6,
			b: 11,
			k: 2,
			r: 3,
		},
		{
			a: 0,
			b: 20,
			k: 5,
			r: 5,
		},
		{
			a: 8,
			b: 40,
			k: 7,
			r: 4,
		},
		{
			a: 15,
			b: 15,
			k: 5,
			r: 1,
		},
		{
			a: 15,
			b: 15,
			k: 7,
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.a, test.b, test.k))
	}
}
