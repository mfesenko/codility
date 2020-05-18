package countdistinctslices

const maxSliceCount = 1000000000

// Solution returns the number of distinct slices in array a.
// If the number of distinct slices is greater than 1,000,000,000, then the function returns 1,000,000,000.
// Lenght of the arrau a is an integer in the range [1..100,000].
// Parameter m is an integer in the range [0..100,000].
// Elements of array a are integers in the range [0..m].
func Solution(m int, a []int) int {
	flags := make([]bool, m+1)
	n := len(a)
	front := 0
	sliceCount := 0
	for back := 0; back < n; back++ {
		for front < n && !flags[a[front]] {
			flags[a[front]] = true
			front++
			sliceCount += front - back
		}
		if sliceCount > maxSliceCount {
			return maxSliceCount
		}
		if front == n {
			break
		}
		flags[a[back]] = false
	}
	return sliceCount
}
