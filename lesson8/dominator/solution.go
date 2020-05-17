package dominator

// Solution returns index of any element of array a in which the dominator of a occurs.
// If the array a does not have a dominator, returns -1.
// The dominator of the array a is the value that occurs in more than half of the elements of a.
// Length of the array a is an integer in the range [0..100,000].
// Elements of the array a are integersin the range [−2,147,483,648..2,147,483,647].
func Solution(a []int) int {
	size := 0
	candidate := 0
	for _, item := range a {
		if size == 0 || item == candidate {
			candidate = item
			size++
			continue
		}
		size--
	}

	count := 0
	lastIndex := -1
	for i, item := range a {
		if item == candidate {
			count++
			lastIndex = i
		}
	}

	if count > len(a)/2 {
		return lastIndex
	}

	return -1
}
