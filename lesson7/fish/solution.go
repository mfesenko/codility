package fish

type stack struct {
	items []int
	size  int
}

// Solution returns the number of fish that will stay alive.
// Array a contains sizes of the fish, array b contains directions of the fish.
// If b[i] == 0, then fish i is flowing upstream, otherwise it's flowing downstream.
// If i < j, then fish i is upstream of fish j. If two fish move in opposite directions
// and there are not other fish between them, they will eventually meet each other.
// Once the two fish meet, the bigger fish eats the smaller fish.
// Arrays a and b have the name length n which is an integer in the range [1..100,000].
// Elements of array a are all distinct integers in the range [0..1,000,000,000].
// Element of array b are integers that can have one of the following values: 0, 1.
func Solution(a []int, b []int) int {
	goingUpstream := newStack()
	goingDownstream := newStack()
	for i, size := range a {
		if b[i] == 1 {
			goingUpstream.add(size)
			continue
		}

		for !goingUpstream.empty() {
			nextObstacle := goingUpstream.peek()
			if nextObstacle > size {
				break
			}
			goingUpstream.pop()
		}

		if goingUpstream.empty() {
			goingDownstream.add(size)
		}
	}
	return goingUpstream.len() + goingDownstream.len()
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) add(a int) {
	s.items = append(s.items, a)
	s.size++
}

func (s *stack) peek() int {
	return s.items[s.size-1]
}

func (s *stack) pop() {
	if s.size > 0 {
		s.size--
		s.items = s.items[:s.size]
	}
}

func (s *stack) empty() bool {
	return s.size == 0
}

func (s *stack) len() int {
	return s.size
}
