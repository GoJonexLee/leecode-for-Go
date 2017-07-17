package medium

/*
问题描述：
中序遍历二叉树

算法描述：递归
 */

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	re := []int{}

	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		f(root.Left)
		re = append(re, root.Val)
		f(root.Right)
	}

	f(root)
	return re
}
