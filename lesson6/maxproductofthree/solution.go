package maxproductofthree

import "sort"

// Solution returns the maximal product of a triplet in array a.
// Length of the array a is an integer in the range [3..100,000].
// Element of the array a are integers in the range [−1,000..1,000].
func Solution(a []int) int {
	sort.Ints(a)
	n := len(a)
	lastItem := a[n-1]
	top := a[n-2] * a[n-3]
	bottom := a[0] * a[1]
	if lastItem > 0 && bottom > top {
		top = bottom
	}
	return lastItem * top
}
