package longestpassword

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
			s: "test 5 a0A pass007 ?xy1",
			r: 7,
		},
		{
			s: "a bb c90 d?a",
			r: -1,
		},
		{
			s: "pass2 perfectPas893 9 shortPas9 cc",
			r: 13,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(test.s))
	}
}
