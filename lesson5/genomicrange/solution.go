package genomicrange

const nucleotidCount = 4

var impactFactors = map[rune]int{
	'A': 1,
	'C': 2,
	'G': 3,
	'T': 4,
}

// Solution returns an array of answers to the queries described by p and q.
// Every k-th query k asks the minimal impact factor of nucleotides contained in the DNA sequence s in the range [p[k]..q[k]].
// DNA sequence s contains only upper-case English letters A, C, G, T.
// Impact factors of nucleotides: A = 1, C = 2, G = 3, T = 4.
// Length n of the DNA sequence s is an integer within the range [1..100,000].
// Length k of the array of queries is an integer within the range [1..50,000].
// Elements of arrays p, q are integers in the range [0..n − 1].
// For each 0 ≤ k < m, p[k] ≤ q[k].
func Solution(s string, p []int, q []int) []int {
	counts := buildCounts(s)
	m := len(p)
	result := make([]int, m)
	for i := 0; i < m; i++ {
		result[i] = findMinNucleotid(counts, p[i], q[i])
	}
	return result
}

func buildCounts(sequence string) [][]int {
	n := len(sequence)
	counts := make([][]int, nucleotidCount)
	for i := 0; i < nucleotidCount; i++ {
		counts[i] = make([]int, n+1)
	}
	for i, nucleotid := range sequence {
		countIndex := impactFactors[nucleotid] - 1
		for j := 0; j < nucleotidCount; j++ {
			prevCount := counts[j][i]
			if j == countIndex {
				prevCount++
			}
			counts[j][i+1] = prevCount
		}
	}
	return counts
}

func findMinNucleotid(counts [][]int, start int, end int) int {
	result := 0
	for i, count := range counts {
		if count[end+1] > count[start] {
			result = i + 1
			break
		}
	}
	return result
}
