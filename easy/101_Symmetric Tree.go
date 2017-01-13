package easy

/*
问题描述：
判断一颗二叉树是否为对称树

算法描述：时间复杂度O(log n), 空间复杂度O(1)
1.定义递归函数f，判断节点的左右子树是否为对称树；
2.递归函数中定义返回状态；
3.递归调用。
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var f func(l, r *TreeNode) bool

	f = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		}

		return (left.Val == right.Val) && f(left.Right, right.Left) && f(left.Left, right.Right)
	}

	return f(root.Left, root.Right)

}
