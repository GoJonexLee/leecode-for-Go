package easy

/*
问题描述：
给定一个整数切片nums以及一个目标值，在nums中找出两个相加等于目标值的整数，并以切片的形式返回这两个整数在nums中的下标

算法描述：时间复杂度O(n),空间复杂度O(n)
遍历nums切片，下标i，整数v，将整数和其对应的下标以k-v的形式存储在map中；
 */

func toSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
    for k, v := range nums {
        if idx, ok := m[target - v]; ok {
            return []int{idx, k}
        }
        m[v] = k
    }
    return nil
}
