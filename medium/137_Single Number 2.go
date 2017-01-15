package medium

/*
问题描述：
给定的整数切片中，除了一个整数出现一次外，其余都出现了三次，找到出现一次的那个整数

算法描述：时间复杂度O(n), 空间复杂度O(1)
 */

func singleNumber(nums []int) int {
	one, two, three := 0, 0, 0
	for _, num := range nums {
		two |= one & num
		one ^= num

		three ^= one & two
		one &= ~three
		two &= ~three
	}
	return one
}
