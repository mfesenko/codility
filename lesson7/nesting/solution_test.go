package nesting

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
			s: "(()(())())",
			r: 1,
		},
		{
			s: "())",
			r: 0,
		},
		{
			s: "",
			r: 1,
		},
		{
			s: "((",
			r: 0,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.s))
	}
}
