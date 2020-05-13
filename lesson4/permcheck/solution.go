package permcheck

// Solution returns 1 if array a is a permutation and 0 if it is not.
// A permutation contains all elements from 1 to n only once.
// Length of the array a is an integer in the range [1..100,000].
// Elements in the array a are integers in the range [1..1,000,000,000].
func Solution(a []int) int {
	if isPermutation(a) {
		return 1
	}
	return 0
}

func isPermutation(a []int) bool {
	n := len(a)
	flags := make([]bool, n+1)
	for _, item := range a {
		if item > n || flags[item] {
			return false
		}
		flags[item] = true
	}
	return true
}
