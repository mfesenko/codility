package tapeequilibrium

// Solution returns a minimal absolute difference between subarrays that can be achieved by splitting
// given array a into two non-empty parts.
// Length of the array a is an integer in the range [2..100,000].
// Elements in the array a are integers in the range [−1,000..1,000].
func Solution(a []int) int {
	leftSum := a[0]
	rightSum := sum(a[1:])
	bestDiff := diff(leftSum, rightSum)
	for i := 1; i < len(a)-1; i++ {
		leftSum += a[i]
		rightSum -= a[i]
		diff := diff(leftSum, rightSum)
		if diff < bestDiff {
			bestDiff = diff
		}
	}
	return bestDiff
}

func sum(a []int) int {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum
}

func diff(a int, b int) int {
	return abs(a - b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
