package oddoccurences

// Solution returns the value of an unpaired item in the array.
// Length of the array a is an odd integer in the range [1..1,000,000].
// Elements in the array a are integers in the range [1..1,000,000,000].
// All values except one in the array a occur an even number of times.
func Solution(a []int) int {
	index := map[int]struct{}{}
	for _, item := range a {
		if _, ok := index[item]; ok {
			delete(index, item)
		} else {
			index[item] = struct{}{}
		}
	}
	for item := range index {
		return item
	}
	return -1
}
