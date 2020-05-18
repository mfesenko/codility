package tieropes

// Solution returns the maximum number of ropes of length greater than or equal
// to k that can be created by tying adjacent ropes from the array a.
// Parameter k is an integer in the range [1..1,000,000,000].
// Length of the array a is an integer in the range [1..100,000].
// Elements of the array a are integers in the range [1..1,000,000,000].
func Solution(k int, a []int) int {
	count := 0
	currentLength := 0
	for _, rope := range a {
		currentLength += rope
		if currentLength >= k {
			count++
			currentLength = 0
		}
	}
	return count
}
