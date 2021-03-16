package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// Time: Avg: O(n); Worst: O(n)
// Space: O(n) where h is the max height of the tree
func NodeDepths(root *BinaryTree) int {
	return traverseNode(root, 0)
}

func traverseNode(node *BinaryTree, depth int) int {

	if node == nil {
		return 0
	}

	return depth + traverseNode(node.Left, depth+1) + traverseNode(node.Right, depth+1)
}
