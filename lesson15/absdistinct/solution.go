package absdistinct

import "sort"

// Solution returns the number of distinct absolute values among the elements of array a.
// Array a is sorted in non-decreasing order.
// Length of the array a is an integer in the range [1..100,000].
// Elements of array a are integers in range [−2,147,483,648..2,147,483,647].
func Solution(a []int) int {
	n := len(a)
	absValues := make([]int, n)
	for i, item := range a {
		absValues[i] = abs(item)
	}
	sort.Ints(absValues)

	front := 0
	distinctCount := 0
	for back := 0; back < n; {
		distinctCount++
		for front < n && absValues[back] == absValues[front] {
			front++
		}

		if front == n {
			break
		}

		back = front
	}
	return distinctCount
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
