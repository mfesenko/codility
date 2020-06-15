package treeheight

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

type test struct {
	s string
	r int
}

func TestSolution(t *testing.T) {
	tests := []test{
		{
			s: "(5, (3, (20, None, None), (21, None, None)), (10, (1, None, None), None))",
			r: 2,
		},
		{
			s: "(5, None, None)",
			r: 0,
		},
		{
			s: "()",
			r: -1,
		},
		{
			s: "(2, (1, None, None), (3, (4, (5, None, None), None)), Nonde)",
			r: 3,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.r, Solution(parseTree(test.s)))
	}
}

func parseTree(s string) *Tree {
	var nodes []*Tree
	var curNode *Tree

	for _, r := range s {
		if r == ')' {
			curNode = nil
			n := len(nodes)
			if n == 0 {
				continue
			}

			nodes[n-3].L = nodes[n-2]
			nodes[n-3].R = nodes[n-1]
			nodes = nodes[:n-2]
		}

		if r == ',' {
			curNode = nil
		}

		if r == 'N' {
			nodes = append(nodes, nil)
		}

		if unicode.IsDigit(r) {
			if curNode == nil {
				curNode = &Tree{}
				nodes = append(nodes, curNode)
			}
			curNode.X = curNode.X*10 + int(r-'0')
		}
	}

	if len(nodes) > 0 {
		return nodes[0]
	}

	return nil
}
