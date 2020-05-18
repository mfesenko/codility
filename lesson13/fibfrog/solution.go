package fibfrog

// Solution returns the minimum number of jumps by which the frog can get to the other side of the river.
// If the frog cannot reach the other side of the river, the function should return −1.
// The frog starts at position -1 and needs to get to the position len(a).
// The frog can jump between the leaves over any distance that is equal to a Fibonacci number.
// Elements of the array a represent positions on the river. If a[i] == 1, then there's a leaf on position i.
// Length of the array a is an integer in the range [0..100,000].
// Elements of the array a are integers that can have one of the following values: 0, 1.
func Solution(a []int) int {
	a = append(a, 1)
	n := len(a)
	jumps := possibleJumps()
	cost := make([]int, n)
	for _, j := range jumps {
		if j > n {
			break
		}
		if a[j-1] == 1 {
			cost[j-1] = 1
		}
	}
	for i := range a {
		if a[i] == 0 || cost[i] > 0 {
			continue
		}
		minCost := n
		found := false
		for _, j := range jumps {
			k := i - j
			if k < 0 {
				break
			}
			prevCost := cost[k]
			if prevCost > 0 && prevCost < minCost {
				minCost = prevCost
				found = true
			}
		}
		if found {
			cost[i] = minCost + 1
		}
	}
	result := cost[n-1]
	if result == 0 {
		return -1
	}
	return result
}

func possibleJumps() []int {
	n := 26
	f := make([]int, n)
	f[1] = 1
	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[2:]
}
