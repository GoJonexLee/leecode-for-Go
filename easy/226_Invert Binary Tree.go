package easy

/*
问题描述：
反转二叉树，即将每个节点的左右子树交换位置

算法描述：时间复杂度O(log n),空间复杂度O(1)
1.交换左右结点；
2.递归。
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}