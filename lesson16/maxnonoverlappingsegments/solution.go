package maxnonoverlappingsegments

// Solution returns the size of largest non-overlapping set of segments
// Each segment i is described by start position a[i] and end b[i] (inclusive).
// Segments are overlapping if they share at least one point.
// The segments are sorted by their ends (b[i] ≤ b[i+1]).
// Length of the arrays a, b is an integer in the range [0..30,000].
// Elements of arrays a, b are integers in the range [0..1,000,000,000].
func Solution(a []int, b []int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	lastEnd := b[0]
	count := 1
	for i := 1; i < n; i++ {
		start := a[i]
		if start > lastEnd {
			lastEnd = b[i]
			count++
		}
	}
	return count
}
