package medium

/*
问题描述：
给定一个长度为n的整数切片，切片中的数据为：0 -- n,并且不重复，求缺少的那个整数
ps:要求时间复杂度O(n), 空间复杂度O(1)

算法描述：时间复杂度O(n), 空间复杂度O(1)
1. 求和整数切片
2. 根据求和公式所得的值，减去1步骤中的和，即为缺失的整数
 */

func missingNumber(nums []int) int {
	length, sum := len(nums), 0
	for i :=0 ; i < length; i++ {
		sum += nums[i]
	}

	return length*(length+1)/2 - sum
}
