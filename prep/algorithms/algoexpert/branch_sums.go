package main

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// Time: O(n) | Space: O(n)
func BranchSums(root *BinaryTree) []int {
	pathSums := []int{}
	calcPathSums(root, 0, &pathSums)

	return pathSums
}

func calcPathSums(node *BinaryTree, parentSum int, leafSums *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*leafSums = append(*leafSums, parentSum+node.Value)
		return
	}

	calcPathSums(node.Left, parentSum+node.Value, leafSums)
	calcPathSums(node.Right, parentSum+node.Value, leafSums)

	return
}
