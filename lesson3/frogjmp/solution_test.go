package frogjmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	x int
	y int
	d int
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			x: 10,
			y: 85,
			d: 30,
			r: 3,
		},
		{
			x: 10,
			y: 10,
			d: 30,
			r: 0,
		},
		{
			x: 10,
			y: 35,
			d: 5,
			r: 5,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.x, test.y, test.d))
	}
}
