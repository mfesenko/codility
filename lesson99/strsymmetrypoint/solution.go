package strsymmetrypoint

// Solution returns the index of a character that the part of the string on the left of the
// character is a reversal of the part of the string to its right.
// If no characters satisfy this requirement, the function returns -1.
// The length of string s is in the range [0..2,000,000].
func Solution(s string) int {
	n := len(s)
	if n%2 == 0 {
		return -1
	}
	mid := n / 2
	for i := 0; i <= mid; i++ {
		if s[mid+i] != s[mid-i] {
			return -1
		}
	}
	return mid
}
