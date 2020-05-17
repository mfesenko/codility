package nesting

// Solution returns 1 if string s is properly nested and 0 otherwise.
// A string s is properly nested if one of the following conditions is true:
// - s is empty;
// - s has the form "(x)" where x is a properly nested string;
// - s has the form "xy" where x and y are properly nested strings.
// Length of the string s is an integer in the range [0..1,000,000].
// String S contains only characters '(' and ')'.
func Solution(s string) int {
	if isProperlyNested(s) {
		return 1
	}
	return 0
}

func isProperlyNested(s string) bool {
	var stack []rune
	size := 0
	for _, r := range s {
		if r == '(' {
			stack = append(stack, r)
			size++
			continue
		}

		if size == 0 {
			return false
		}

		size--
		stack = stack[:size]
	}
	return size == 0
}
