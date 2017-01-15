package medium

/*
问题描述：
二叉树的前序遍历

算法描述：递归
 */

func preorderTraversal(root *TreeNode) []int {
	re := []int{}

	var f func(*TreeNode, *[]int)
	f = func(root *TreeNode, re *[]int) {
		if root == nil {
			return
		}
		*re = append(*re, root.Val)
		f(root.Left, re)
		f(root.Right, re)
	}

	f(root, &re)
	return re
}
