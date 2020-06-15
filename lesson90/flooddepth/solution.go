package flooddepth

// Solution returns the maximum depth of the water after a landscape is completely
// flooded. The landscape is divided into blocks and described by array a that
// contains the altitude of a block.
// Length of the array a is an integer in the range [1..100,000].
// Elements of the array a are integers in the range [1..100,000,000].
func Solution(a []int) int {
	left := calculateLeftMax(a)
	right := calculateRightMax(a)
	maxDepth := 0
	for i := 1; i < len(a)-1; i++ {
		leftMax := left[i]
		rightMax := right[i]
		height := a[i]
		if height < leftMax && height < rightMax {
			maxDepth = max(maxDepth, min(leftMax, rightMax)-height)
		}
	}
	return maxDepth
}

func calculateLeftMax(a []int) []int {
	n := len(a)
	left := make([]int, n)
	curMax := a[0]
	for i := 1; i < n; i++ {
		left[i] = curMax
		curMax = max(curMax, a[i])
	}
	return left
}

func calculateRightMax(a []int) []int {
	n := len(a)
	right := make([]int, n)
	curMax := a[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = curMax
		curMax = max(curMax, a[i])
	}
	return right
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
