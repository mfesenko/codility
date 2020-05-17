package commonprimedivisors

// Solution returns the number of positions i for which the prime divisors
// of a[i] and b[i] are exactly the same.
// The length of the arrays a and b is an integer in the range [1..6,000].
// Elements of the arrays a, b are integers in the range [1..2,147,483,647].
func Solution(a []int, b []int) int {
	count := 0
	for i := range a {
		if samePrimeDivisors(a[i], b[i]) && samePrimeDivisors(b[i], a[i]) {
			count++
		}
	}
	return count
}

func samePrimeDivisors(a int, b int) bool {
	if b == 1 {
		return true
	}

	gcd := gcd(a, b)
	if gcd == 1 {
		return false
	}

	return samePrimeDivisors(a, b/gcd)
}

func gcd(a int, b int) int {
	return gcdEfficient(a, b, 1)
}

func gcdEfficient(a int, b int, res int) int {
	if a == b {
		return a * res
	}
	aEven := a%2 == 0
	bEven := b%2 == 0
	if aEven && bEven {
		return gcdEfficient(a>>1, b>>1, res<<1)
	}
	if aEven {
		return gcdEfficient(a>>1, b, res)
	}
	if bEven {
		return gcdEfficient(a, b>>1, res)
	}
	if a > b {
		return gcdEfficient(a-b, b, res)
	}
	return gcdEfficient(a, b-a, res)
}
