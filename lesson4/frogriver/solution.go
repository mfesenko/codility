package frogriver

const never = -1

// Solution returns the earliest time when the frog can jump from position 0 to position x+1.
// If the frog is never going to be able to reach position x+1 - returns -1.
// Each element a[i] represent a position where a leaf falls at time i.
// The frog can cross only when leaves appear at every position across the river from 1 to x.
// Length of the array a and position x are integers in the range [1..100,000].
// Elements in the array a are integers in the range [1..x].
func Solution(X int, A []int) int {
	firstTime := make([]int, X)
	for i, a := range A {
		if firstTime[a-1] == 0 {
			firstTime[a-1] = i + 1
		}
	}
	earliestTime := never
	for _, t := range firstTime {
		if t == 0 {
			earliestTime = never
			break
		}
		earliestTime = max(earliestTime, t-1)
	}
	return earliestTime
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
