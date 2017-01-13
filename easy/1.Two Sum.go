package easy

/*
问题描述：
给定一个整数切片nums以及一个目标值，在nums中找出两个相加等于目标值的整数，并以切片的形式返回这两个整数在nums中的下标

算法描述：时间复杂度O(n),空间复杂度O(n)
1.遍历nums切片，下标i，整数v，将整数和其对应的下标以k-v的形式存储在map中；
2.重新遍历nums，下标i，整数v，并在map中查找是否存在（target-v）的整数；
 */

func twoSum(nums []int, target int) []int {
	re, mp := []int{}, make(map[int]int)
	for i, v := range nums {
		mp[v] = i
	}

	for i, v := range nums {
		if index, ok := mp[target - v]; index != i && ok {
			re = append(re, i, index)
			break
		}
	}
	return re
}
