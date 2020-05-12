package binarygap

import (
	"strconv"
	"strings"
)

// Solution returns the length of the longest sequence of zeros surrounded by ones in the binary representation of N.
// N is an integer in the range [1..2,147,483,647].
func Solution(N int) int {
	binary := strconv.FormatInt(int64(N), 2)
	gaps := strings.Split(strings.Trim(binary, "0"), "1")
	maxGapSize := 0
	for _, gap := range gaps {
		maxGapSize = max(maxGapSize, len(gap))
	}
	return maxGapSize
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
