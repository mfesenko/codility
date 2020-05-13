package maxcounters

// Solution returns the value of n counters after applying operations from a.
// The counters support following operations:
// - increase(x) - counter x is increased by 1
// - maxCounter() - all counters are set to the vakue of the maximum counter
// Array a contains m operations that are applied consequtively:
// - if a[i] == n+1 then the operation is maxCounter()
// - if 1 <= a[i] <= n then the operation is increase(a[i])
// n and m are integers in the range [1..100,000].
// Elements in the array a are integer values in the range [1..n+1]
func Solution(n int, a []int) []int {
	counters := make([]int, n)
	offset := 0
	maxCounter := 0
	for _, operation := range a {
		if operation == n+1 {
			offset = maxCounter
			continue
		}
		maxCounter = increaseCounter(counters, operation-1, maxCounter, offset)
	}
	for i := 0; i < n; i++ {
		c := counters[i]
		if c < offset {
			counters[i] = offset
		}
	}
	return counters
}

func increaseCounter(counters []int, i int, max int, offset int) int {
	counter := counters[i]
	if counter < offset {
		counter = offset
	}
	counter++
	counters[i] = counter
	if counter > max {
		return counter
	}
	return max
}
