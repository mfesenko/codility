package countsemiprimes

// Solution returns an array containing consecutive answers to all the queries.
// Query i requires you to find the number of semiprimes in the range (p[i], q[i]), where 1 ≤ p[i] ≤ q[i] ≤ n.
// A semiprime is a natural number that is the product of two prime numbers.
// Parameter n is an integer in the range [1..50,000].
// Arrays p and q have the same length m in the range [1..30,000].
// Elements of arrays p, q are integers in the range [1..n].
func Solution(n int, p []int, q []int) []int {
	f := buildFactorizationTable(n)
	semiprimes := allSemiprimes(n, f)
	sums := prefixSum(semiprimes)
	result := make([]int, len(p))
	for i, start := range p {
		end := q[i]
		result[i] = sums[end+1] - sums[start]
	}
	return result
}

func buildFactorizationTable(n int) []int {
	f := make([]int, n+1)
	for i := 2; i*i <= n; i++ {
		if f[i] > 0 {
			continue
		}

		for j := i * i; j <= n; j += i {
			if f[j] == 0 {
				f[j] = i
			}
		}
	}
	return f
}

func allSemiprimes(N int, f []int) []bool {
	semiprimes := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if len(factorize(f, i)) == 2 {
			semiprimes[i] = true
		}
	}
	return semiprimes
}

func factorize(f []int, i int) []int {
	var factors []int
	for f[i] > 0 {
		factors = append(factors, f[i])
		i = i / f[i]
	}
	factors = append(factors, f[i])
	return factors
}

func prefixSum(semiprimes []bool) []int {
	sums := make([]int, len(semiprimes)+1)
	for i, s := range semiprimes {
		sum := sums[i]
		if s {
			sum++
		}
		sums[i+1] = sum
	}
	return sums
}
