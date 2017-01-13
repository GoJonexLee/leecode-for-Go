package easy

/*
问题描述：
删除整数切片nums中，所有等于val的整数

算法描述：时间复杂度O(n),空间复杂度O(1)
1.re用于记录nums中不等于val的个数；
2.遍历nums，当v！=val时，将v移动到下标为re的位置中；
3.设置最终切片的位置，并返回。
 */

func removeElement(nums []int, val int) int {
	re := 0
	for _, v := range nums {
		if v != val {
			nums[re] = v
			re++
		}
	}
	nums = nums[:re]
	return re
}
