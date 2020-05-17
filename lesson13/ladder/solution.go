package ladder

// Solution returns an array containing the number of different ways
// to climb the ladder with a[i] rungs modulo 2^b[i].
// When climbing the ladder you start on rung 0. When you are standing
// on rung i, you can move to rungs i+1 or i+2.
// The length of arrays a and b is an integer in the range [1..50,000].
// Elements of array a are integers in the range [1..len(a)].
// Element of array b are integers in the range [1..30].
func Solution(A []int, B []int) []int {
	maxRung := findMax(A)
	counts := countWays(maxRung)
	result := make([]int, len(A))
	for i, a := range A {
		result[i] = int(counts[a] % (1 << uint(B[i])))
	}
	return result
}

func findMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func countWays(n int) []uint64 {
	count := make([]uint64, n+2)
	count[1] = 1
	count[2] = 1
	for i := 1; i < n; i++ {
		count[i+1] += count[i]
		count[i+2] += count[i]
	}
	return count
}
