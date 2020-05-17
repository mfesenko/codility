package distinct

import "sort"

// Solution returns the number of distinct values in array a.
// Length of the array a is an integer in the range [0..100,000].
// Elements in the array a are integers in the range [−1,000,000..1,000,000].
func Solution(a []int) int {
	sort.Ints(a)
	count := 0
	prev := 0
	for i, item := range a {
		if i == 0 || item != prev {
			count++
		}
		prev = item
	}
	return count
}
