package triangle

import "sort"

// Solution returns 1 if three elements from array a can form a triangle, otherwise returns 0.
// Length of the array a is an integer in the range [0..100,000].
// Elements of the array a are integers in the range [−2,147,483,648..2,147,483,647].
func Solution(a []int) int {
	if triangularExists(a) {
		return 1
	}
	return 0
}

func triangularExists(a []int) bool {
	sort.Ints(a)
	n := len(a)
	for i := 0; i < n-2; i++ {
		if a[i+2] < a[i+1]+a[i] {
			return true
		}
	}
	return false
}
