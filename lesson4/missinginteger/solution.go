package missinginteger

const maxValue = 1000000

// Solution returns the smallest integer greater than 0 that does not occur in the array a.
// Length of the array a is an integer in the range [1..100,000].
// Elements in the array a are integers in the range [−1,000,000..1,000,000].
func Solution(a []int) int {
	flags := make([]bool, maxValue)
	for _, item := range a {
		if item > 0 {
			flags[item-1] = true
		}
	}
	r := maxValue + 1
	for i, f := range flags {
		if !f {
			r = i + 1
			break
		}
	}
	return r
}
