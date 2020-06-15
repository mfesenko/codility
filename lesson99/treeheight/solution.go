package treeheight

// Tree represents a node of a binary tree
type Tree struct {
	X int
	L *Tree
	R *Tree
}

// Solution returns the height of the binary tree.
// The amount of nodes in the tree is in the range [1..1,000].
// The height of the tree is in the range [0..500].
func Solution(root *Tree) int {
	if root == nil {
		return -1
	}
	return max(Solution(root.L), Solution(root.R)) + 1
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
