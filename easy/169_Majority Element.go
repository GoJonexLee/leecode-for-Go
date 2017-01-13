package easy

/*
问题描述：
求出整数切片中的众数，众数的出现次数大于切片长度的一半

算法描述：时间复杂度O(n), 空间复杂度O(1)
1. 摩尔投票法；
2. 位操作；
 */

func majorityElement(nums []int) int {
	re, ct := 0, 0
	for _, num := range nums {
		if ct == 0 {
			re = num
			ct++
		} else if num == re {
			ct++
		} else {
			ct--
		}
	}
	return re
}
