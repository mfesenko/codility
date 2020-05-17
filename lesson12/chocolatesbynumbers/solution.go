package chocolatesbynumbers

// Solution returns the number of chocolates that you can eat following the rules below.
// Parameter n represents the number of chocolates arranged in a circle, numbered from 0 to n − 1.
// You start to eat the chocolates from number 0. Then you skip m-1 positions in the circle and eat the next one.
// You stop eating the chocolated when you move to a position with an empty wrapper.
// Parameters n and m are integers in the range [1..1,000,000,000].
func Solution(n int, m int) int {
	return n / gcd(n, m)
}

func gcd(a, b int) int {
	return gcdEfficient(a, b, 1)
}

func gcdEfficient(a, b, res int) int {
	if a == b {
		return res * a
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
