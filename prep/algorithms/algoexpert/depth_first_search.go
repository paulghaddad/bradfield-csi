package main

type Node struct {
	Name     string
	Children []*Node
}

// Time: O(v + e) | Space: O(v)
func (n *Node) DepthFirstSearch(array []string) []string {
	namesSeen := make([]string, 0)
	stack := []Node{*n}

	for len(stack) > 0 {
		stackLen := len(stack)
		poppedNode := stack[stackLen-1]
		stack = stack[:stackLen-1]

		namesSeen = append(namesSeen, poppedNode.Name)

		// go in reverse to traverse L to R
		for i := len(poppedNode.Children) - 1; i >= 0; i-- {
			stack = append(stack, *poppedNode.Children[i])
		}
	}

	return namesSeen
}
