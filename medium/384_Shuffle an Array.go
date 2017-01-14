package medium

import (
	"math/rand"
	"time"
)


/*
问题描述：
给定一个整数切片，根据Reset和Shuffle方法返回不同的整数切片，
Reset：原始切片
Shuffle：随机切片

算法描述：时间复杂度O(n), 空间复杂度O(n)
1. Solution中必须记录原始切片；
2. 通过随机数调整切片中数据的位置；
 */

type Solution struct {
	original []int	// 记录原始切片
	rd *rand.Rand	// 设置随机因子
	size int	// 原始切片长度，该字段可通过len()函数替代
}

func Constructor(nums []int) Solution {
	return Solution{original: nums, rd: rand.New(rand.NewSource(time.Now().UnixNano())), size: len(nums)}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	re := make([]int, this.size)
	copy(re, this.original)

	for i := 0; i < this.size; i++ {
		j := this.rd.Intn(this.size)
		re[i], re[j] = re[j], re[i]
	}
	return re
}
