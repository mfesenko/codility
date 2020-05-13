package minavgtwoslice

// Solution returns the starting position of the first slice in the array a with the minimal average.
// Slice contains at least two elements.
// Length of the array a is an integer in the range [2..100,000].
// Elements in the array a are integers in the range [−10,000..10,000].
func Solution(a []int) int {
	sum := prefixSum(a)
	n := len(a)
	minAvg := avg(sum, 0, 1)
	minIndex := 0
	for i := 0; i < n-1; i++ {
		minAvg, minIndex = checkSlice(sum, i, i+1, minAvg, minIndex)
		if i != n-2 {
			minAvg, minIndex = checkSlice(sum, i, i+2, minAvg, minIndex)
		}
	}
	return minIndex
}

func prefixSum(a []int) []int {
	sum := make([]int, len(a)+1)
	for i, item := range a {
		sum[i+1] = sum[i] + item
	}
	return sum
}

func checkSlice(sum []int, start int, end int, minAvg float32, minIndex int) (float32, int) {
	curAvg := avg(sum, start, end)
	if curAvg < minAvg {
		return curAvg, start
	}
	return minAvg, minIndex
}

func avg(sum []int, start int, end int) float32 {
	return float32(sum[end+1]-sum[start]) / float32(end-start+1)
}
