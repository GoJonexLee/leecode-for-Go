package easy

/*
问题描述：
判断一颗二叉树是否为平衡二叉树

算法描述：时间复杂度O(n log n), 空间复杂度O(1)
1.定义深度函数deep，求出节点的最大深度；
2.递归判断各个节点是否为平衡二叉树

注意：目前的算法存在重复计算
优化方法：动态规划制表法，时间复杂度O(log n), 空间复杂度O(n)
 */

func deep(rt *TreeNode) int {
	if rt == nil {
		return 0
	}

	if rt.Left == nil && rt.Right == nil {
		return 1
	}

	if ld, rd := deep(rt.Left), deep(rt.Right); ld >= rd {
		return 1 + ld
	} else {
		return 1 + rd
	}
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if lDep, rDep := deep(root.Left), deep(root.Right); lDep - rDep > 1 || lDep - rDep < -1 {
		return false
	}

	return isBalanced(root.Right) && isBalanced(root.Left)
}
