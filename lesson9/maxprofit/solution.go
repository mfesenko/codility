package maxprofit

// Solution returns the maximum possible profit from one transaction during the period
// described by array a that contains daily prices of a stock share.
// If it was impossible to gain any profit, it returns 0.
// Length of the array is an integer in the range [0..400,000].
// Elements of the array a are integers in the range [0..200,000].
func Solution(a []int) int {
	maxGain := 0
	minPrice := 0
	for i, price := range a {
		if i == 0 {
			minPrice = price
			continue
		}

		maxGain = max(maxGain, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxGain
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
