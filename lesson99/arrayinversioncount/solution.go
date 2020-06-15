package arrayinversioncount

const maxInversionCount = 1000000000

// Solution returns the number of intervions in array a.
// If the number of inversion exceeds 1,000,000,000, it returns -1.
// Length of the array a is an integer in the range [0..100,000].
// Elements of array a are integers in the range [−2,147,483,648..2,147,483,647].
func Solution(a []int) int {
	c := copySlice(a)
	count := mergeSort(c, 0, len(c)-1)
	if count > maxInversionCount {
		return -1
	}
	return count
}

func copySlice(a []int) []int {
	c := make([]int, len(a))
	copy(c, a)
	return c
}

func mergeSort(a []int, start int, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)/2
	count := mergeSort(a, start, mid)
	count += mergeSort(a, mid+1, end)
	count += merge(a, start, mid, end)
	return count
}

func merge(a []int, start int, mid int, end int) int {
	count := 0
	n := mid - start + 1
	m := end - mid
	left := copySlice(a[start : mid+1])
	right := copySlice(a[mid+1 : end+1])

	i := 0
	j := 0
	k := start

	for i < n && j < m {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
			count += mid + 1 - start - i
		}
		k++
	}

	for i < n {
		a[k] = left[i]
		i++
		k++
	}

	for j < m {
		a[k] = right[j]
		j++
		k++
	}

	return count
}
