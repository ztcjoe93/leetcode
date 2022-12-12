package main

func sortedArrayToBST(nums []int) *TreeNode {
	depth := len(nums) / 2
	medianNode := TreeNode{Val: nums[depth]}

	if len(nums[0:depth]) > 0 {
		medianNode.Left = sortedArrayToBST(nums[0:depth])
	}

	if len(nums[depth+1:]) > 0 {
		medianNode.Right = sortedArrayToBST(nums[depth+1:])
	}

	return &medianNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
