package hard

/*
问题描述：
二叉树的后续遍历

算法描述：时间复杂度O(log n), 空间复杂度O(n)

坑：
f函数的第二个参数一定要传切片的指针，否则会发生切片找不到的效果
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	result := []int{}

	var dp func(treeNode *TreeNode)

	dp = func(root *TreeNode) {
		if root == nil {
			return
		}
		dp(root.Left)
		dp(root.Right)
		result = append(result, root.Val)
	}
	dp(root)
	return result
}
