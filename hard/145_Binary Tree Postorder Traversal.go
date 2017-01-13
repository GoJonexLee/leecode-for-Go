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

	var f func(*TreeNode, *[]int)
	f = func(rt *TreeNode, rl *[]int) {
		if rt == nil {
			return
		}

		if rt.Left == nil && rt.Right == nil {
			*rl = append(*rl, rt.Val)
			return
		}

		f(rt.Left, rl)
		f(rt.Right, rl)
		*rl = append(*rl, rt.Val)
	}

	re := []int{}
	f(root, &re)
	return re
}
