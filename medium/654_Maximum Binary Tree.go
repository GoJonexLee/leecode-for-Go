package medium

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	
	if len(nums) == 0 {
		return nil
	}

	maxIdx := maxIndex(nums)
	tree := new(TreeNode)
	tree.Val = nums[maxIdx]
	tree.Left = constructMaximumBinaryTree(nums[: maxIdx])
	tree.Right = constructMaximumBinaryTree(nums[maxIdx+1 :])
	
	return tree
}

func maxIndex(nums []int) int {
	maxIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}