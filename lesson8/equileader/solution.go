package equileader

// Solution returns the number of equi leaders in array a.
// The leader of the array a is the value that occurs in more than half of the elements of a.
// The equi leader of the array a is an index i such that 0 ≤ i < len(a) − 1 where two sequences
// a[0], s[1], ..., s[i] and a[i + 1], a[i + 2], ..., a[len(a) − 1] have the same leaders.
// Length of the array a is an integer in the range [1..100,000].
// Elements of the array a are integers in the range [−1,000,000,000..1,000,000,000].
func Solution(a []int) int {
	candidate := findLeaderCandidate(a)
	n := len(a)
	totalCount := 0
	runningCount := make([]int, n)
	for i, item := range a {
		if candidate == item {
			totalCount++
		}
		runningCount[i] = totalCount
	}

	if totalCount <= n/2 {
		return 0
	}

	equiLeaderCount := 0
	for i := 0; i < n-1; i++ {
		leftCount := runningCount[i]
		rightCount := totalCount - leftCount
		if leftCount > (i+1)/2 && rightCount > (n-i-1)/2 {
			equiLeaderCount++
		}
	}
	return equiLeaderCount
}

func findLeaderCandidate(a []int) int {
	size := 0
	value := 0

	for _, item := range a {
		if size == 0 || value == item {
			value = item
			size++
			continue
		}
		size--
	}

	return value
}
