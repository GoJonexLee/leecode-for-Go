package easy

/*
算法描述：
整数切片中，除一个整数只出现1次外，其余均出现2次，找出只出现一次的那个整数

算法描述：时间复杂度O(n), 空间复杂度O(1)
1.两个相同的数或字符串，异或操作为0；
2.遍历nums切片，异或最终的结果即为该整数。
 */

func singleNumber(nums []int) int {
	tmp := 0
	for _, num := range nums {
		tmp ^= num
	}
	return tmp
}
