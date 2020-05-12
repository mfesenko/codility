package cyclicrotation

// Solution returns the result of rotating the array a k times.
// Length of the array a and k are integers in the range [0..100].
// Elements in the array a are integers in the range [−1,000..1,000].
func Solution(a []int, k int) []int {
	size := len(a)
	if size < 2 {
		return a
	}
	index := size - k%size
	return append(a[index:], a[:index]...)
}
