package numbersolitaire

const maxStepCount = 6

// Solution returns the maximal result of the game that can be achieved on the board represented by array a.
// The goal of the game is to move the pebble from square number 0 to square number len(a) − 1.
// During each turn the pebble can me moved from the current position i to positions in range [i+1...i+6].
// We mark each square the peddle visits.
// The result of the game is the sum of the numbers written on all marked squares.
// Length of the array a is an integer in the range [2..100,000].
// Elements of the array a are integers in the range [−10,000..10,000].
func Solution(a []int) int {
	n := len(a)
	memo := make([]int, len(a))
	for i, item := range a {
		if i == 0 {
			memo[i] = item
			continue
		}
		end := max(0, i-maxStepCount)
		score := memo[i-1]
		for j := i - 2; j >= end; j-- {
			score = max(memo[j], score)
		}
		memo[i] = score + item
	}
	return memo[n-1]
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
