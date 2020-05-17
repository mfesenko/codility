package stonewall

type stack struct {
	items []int
	size  int
}

// Solution returns the minimum number of blocks needed to build the wall described by array h.
// Each element h[i] contains the height of the wall from i to i+1 meters from the start.
// Length of the array h is an integer in the range [1..100,000].
// Elements of the array h are integers in the range [1..1,000,000,000].
func Solution(h []int) int {
	completeBlocks := newStack()
	incompleteBlocks := newStack()

	for _, height := range h {
		for incompleteBlocks.len() > 0 {
			prevHeight := incompleteBlocks.peek()
			if height < prevHeight {
				incompleteBlocks.pop()
				completeBlocks.add(prevHeight)
				continue
			}
			if height > prevHeight {
				incompleteBlocks.add(height)
			}
			break
		}

		if incompleteBlocks.len() == 0 {
			incompleteBlocks.add(height)
			continue
		}
	}

	return completeBlocks.len() + incompleteBlocks.len()
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) add(item int) {
	s.items = append(s.items, item)
	s.size++
}

func (s *stack) peek() int {
	return s.items[s.size-1]
}

func (s *stack) pop() {
	s.size--
	s.items = s.items[:s.size]
}

func (s *stack) len() int {
	return s.size
}
