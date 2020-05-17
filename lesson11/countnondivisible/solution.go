package countnondivisible

// Solution returns a sequence of integers representing the amount of non-divisors in the array a.
// Length of the array a is an integer in the range [1..50,000].
// Elements of the array a are integers in the range [1..2 * len(a)].
func Solution(a []int) []int {
	n := len(a)
	maxValue := n << 1
	divisors := buildDivisors(maxValue)
	counts := count(maxValue, a)
	result := make([]int, n)
	cache := make([]int, maxValue+1)
	for i, item := range a {
		cached := cache[item]
		if cached == 0 {
			count := 0
			for _, d := range divisors[item] {
				count += counts[d]
			}
			cached = n - count
			cache[item] = cached
		}
		result[i] = cached
	}
	return result
}

func count(maxValue int, a []int) []int {
	counts := make([]int, maxValue+1)
	for _, item := range a {
		counts[item]++
	}
	return counts
}

func buildDivisors(n int) [][]int {
	divisors := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		d := make([]int, 0, 8)
		d = append(d, 1)
		if i != 1 {
			d = append(d, i)
		}
		divisors[i] = d
	}

	for i := 2; i <= n; i++ {
		for j := i << 1; j <= n; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	return divisors
}
