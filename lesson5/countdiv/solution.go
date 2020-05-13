package countdiv

// Solution returns the number of integers in the range [a..b] that are divisible by k.
// Parameters a and b are integers in the range [0..2,000,000,000], a ≤ b.
// Parameter k is an integer within the range [1..2,000,000,000].
func Solution(a int, b int, k int) int {
	first := firstDivisible(a, k)

	if first > b {
		return 0
	}

	interval := b - first
	return interval/k + 1
}

func firstDivisible(a int, k int) int {
	first := (a / k) * k
	if a%k != 0 {
		first += k
	}
	return first
}
