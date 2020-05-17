package maxslicesum

// Solution returns the maximum sum of any slice of array a.
// Length of the array a is an integer in the range [1..1,000,000].
// Elements of the array a are integers in the range [−1,000,000..1,000,000].
// The result is an integer in the range [−2,147,483,648..2,147,483,647].
func Solution(A []int) int {
	curSum := 0
	maxSum := -1 << 31
	for _, a := range A {
		curSum = max(a, a+curSum)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
