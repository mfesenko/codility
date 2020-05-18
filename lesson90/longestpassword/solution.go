package longestpassword

import (
	"strings"
	"unicode"
)

// Solution returns the length of the longest word from string s that is a valid password.
// If there are no valid passwords in string s, then returns -1.
// A word needs to comply with the following rules to be a valid password:
// - contains only alphanumerical characters (a-z, A-Z, 0-9)
// - there is an even number of letters
// - there is an odd number of digits
// Length of the string s is an inteter in the range [1..200]
// String s contains only printable ASCII characters and spaces
func Solution(s string) int {
	words := strings.Split(s, " ")
	maxLen := -1
	for _, word := range words {
		if isValidPassword(word) {
			maxLen = max(maxLen, len(word))
		}
	}
	return maxLen
}

func isValidPassword(word string) bool {
	n := len(word)
	if n%2 == 0 {
		return false
	}
	digitCount := 0
	for _, r := range word {
		if unicode.IsLetter(r) {
			continue
		}
		if unicode.IsDigit(r) {
			digitCount++
			continue
		}
		return false
	}
	return (digitCount % 2) == 1
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
