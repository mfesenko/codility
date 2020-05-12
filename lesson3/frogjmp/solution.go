package frogjmp

// Solution returns the minimal number of jumps the frog needs to make
// to move from position x to a position greater than or equal to y.
// The frog jumps with a fixed distance d.
// x, y and d are integers in the range [1..1,000,000,000], x ≤ y.
func Solution(x int, y int, d int) int {
	distance := y - x
	jumps := distance / d
	if distance%d != 0 {
		jumps++
	}
	return jumps
}
