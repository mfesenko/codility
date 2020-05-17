package ladder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	a := []int{4, 4, 5, 5, 1}
	b := []int{3, 2, 4, 3, 1}
	r := []int{5, 1, 8, 0, 1}
	assert.Equal(t, r, Solution(a, b))
}
