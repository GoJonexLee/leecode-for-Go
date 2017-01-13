package easy

/*
问题描述：
在二叉树中，从根节点到叶子节点的路径中，是否存在节点值的和为sum的路径

算法描述：时间复杂度O(log n), 空间复杂度O(1)
1.函数结束状态；
2.左右递归。
 */

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	leave := sum - root.Val
	return hasPathSum(root.Left, leave) || hasPathSum(root.Right, leave)
}
