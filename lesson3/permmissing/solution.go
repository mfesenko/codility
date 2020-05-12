package permmissing

// Solution returns the value of the missing element in the array a.
// The array a contains n distinct elements from the range [1..(n + 1)].
// Length of the array is an integer in the range [0..100,000].
func Solution(a []int) int {
	n := len(a) + 1
	index := make([]bool, n)
	for _, item := range a {
		index[item-1] = true
	}
	missing := -1
	for i, ok := range index {
		if !ok {
			missing = i + 1
			break
		}
	}
	return missing
}
