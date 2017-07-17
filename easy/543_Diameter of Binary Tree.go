package easy

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var re int

	deep(root, &re)
	return re
}

func deep(node *TreeNode, dp *int) int {
	if node == nil {
		return 0
	}

	left, right := deep(node.Left, dp), deep(node.Right, dp)
	if *dp < left + right {
		*dp = left + right
	}

	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}