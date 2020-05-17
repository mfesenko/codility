package brackets

// Solution returns 1 if string s is properly nested and 0 otherwise.
// A string s is properly nested if one of the following conditions is true:
// - s is empty;
// - s has the form "(x)" or "[x]" or "{x}" where t is a properly nested string;
// - s has the form "xy" where x and y are properly nested strings.
// Length of the string s is an integer in the range [0..200,000].
// String s contains only following characters: '(', '{', '[', ']', '}', ')'.
func Solution(s string) int {
	if isProperlyNested(s) {
		return 1
	}
	return 0
}

func isProperlyNested(s string) bool {
	brackets := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	var stack []rune
	stackSize := 0
	for _, r := range s {
		expectedPrev, isClosing := brackets[r]
		if !isClosing {
			stack = append(stack, r)
			stackSize++
			continue
		}

		if stackSize == 0 {
			return false
		}

		prev := stack[stackSize-1]
		if prev != expectedPrev {
			return false
		}

		stackSize--
		stack = stack[:stackSize]
	}

	return stackSize == 0
}
