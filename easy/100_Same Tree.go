package easy

/*
问题描述：
判断两颗二叉树是否相同

算法描述：时间复杂度O(log n), 空间复杂度O(1)
1.判断两棵树的节点值是否相等；
2.递归判断两棵树的左子树和右子数；
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p ==nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}

}
