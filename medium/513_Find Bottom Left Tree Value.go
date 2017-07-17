package medium

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result, maxDeep := root.Val, 1
	deep(root, 1, &maxDeep, &result)

	return result
}

func deep(tree *TreeNode, dp int, maxDeep, val *int) {
	if tree == nil {
		return
	}

	if dp > *maxDeep {
		*maxDeep = dp
		*val = tree.Val
	}
	dp += 1
	deep(tree.Left, dp, maxDeep, val)
	deep(tree.Right, dp, maxDeep, val)
}