package binarygap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := map[int]int{
		4:    0, //         100
		9:    2, //        1001
		15:   0, //        1111
		20:   1, //       10100
		529:  4, //  1000010001
		1041: 5, // 10000010001
	}
	for n, r := range tests {
		assert.Equal(t, r, Solution(n))
	}
}
