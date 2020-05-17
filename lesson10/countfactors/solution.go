package countfactors

// Solution returns the number of factors of a number n.
// Parameter n is an integer in the range [1..2,147,483,647].
func Solution(n int) int {
	factorCount := 0
	i := 1
	for i*i < n {
		if n%i == 0 {
			factorCount += 2
		}
		i++
	}
	if i*i == n {
		factorCount++
	}
	return factorCount
}
