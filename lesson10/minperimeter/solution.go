package minperimeter

// Solution returns the minimal perimeter of any rectangle with area exactly equal to n.
// Parameter n is an integer in the range [1..1,000,000,000].
func Solution(n int) int {
	divisors := findDivisors(n)
	minPerimeter := 0
	for i, a := range divisors {
		curPerimeter := perimeter(a, n/a)
		if i == 0 || curPerimeter < minPerimeter {
			minPerimeter = curPerimeter
		}
	}
	return minPerimeter
}

func findDivisors(N int) []int {
	var divisors []int
	i := 1
	for i*i < N {
		if N%i == 0 {
			divisors = append(divisors, i)
		}
		i++
	}
	if i*i == N {
		divisors = append(divisors, i)
	}
	return divisors
}

func perimeter(a, b int) int {
	return 2 * (a + b)
}
