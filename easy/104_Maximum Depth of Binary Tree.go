package easy

/*
问题描述：
求二叉树的深度

算法描述：时间复杂度O(log n),空间复杂度O(1)
1.定义递归函数的返回状态；
2.分别求二叉树左右子树的最大深度；
3.返回左右子树的深度最大值，并且+1。
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if ml, mr := maxDepth(root.Left), maxDepth(root.Right); ml > mr {
		return 1 + ml
	} else {
		return 1 + mr
	}
}
