package strsymmetrypoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	s string
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			s: "racecar",
			r: 3,
		},
		{
			s: "x",
			r: 0,
		},
		{
			s: "zxc",
			r: -1,
		},
		{
			s: "boom",
			r: -1,
		},
		{
			s: "qwwq",
			r: -1,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.s))
	}
}
