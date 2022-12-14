package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	var leftDepth int
	var rightDepth int

	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}

	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}

	if leftDepth > rightDepth {
		depth += leftDepth
	} else {
		depth += rightDepth
	}

	return depth
}
