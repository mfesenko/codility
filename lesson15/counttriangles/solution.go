package counttriangles

import "sort"

// Solution returns the number of triangular triplets in this array.
// A triplet (i, j, k) is triangular if it is possible to build a triangle with sides of length a[i], a[j], a[k].
// Length of the array a is an integer in the range [0..1,000].
// Elements of the array a are integers in the range [1..1,000,000,000].
func Solution(a []int) int {
	sort.Ints(a)
	n := len(a)
	count := 0
	for i := 0; i < n-2; i++ {
		j := i + 2
		for k := i + 1; k < n-1; k++ {
			for j < n && a[i]+a[k] > a[j] {
				j++
			}
			count += j - k - 1
		}
	}
	return count
}
