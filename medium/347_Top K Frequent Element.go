package medium

import (
	"sort"
)

/*
问题描述：
给定一个整数切片，求在该切片中出现次数最多的K个，并以切片的形势返回
ps:假设K肯定有效，例如1<= k <= len(nums);
   时间复杂度小于O(n log n)

算法描述：时间复杂度O(n), 空间复杂度O(n)
1.通过map统计切片中整数的频率；
2.遍历map，记录topK

mp:统计nums中的频次
re：函数返回的结果

mmp：反向统计频次的整数切片
tp：频次排序
LEE标签：统计K个整数
 */

func topKFrequent(nums []int, k int) []int {
	mp, re := make(map[int]int), make([]int, 0, k)
	for _, num := range nums {
		mp[num]++
	}

	mmp := make(map[int][]int)
	for k, v := range mp {
		mmp[v] = append(mmp[v], k)
	}

	tp := []int{}
	for k, _ := range mmp {
		tp = append(tp, k)
	}

	sort.Ints(tp)
	LEE:
	for i := len(tp)-1; i >= 0; i-- {
		for _, num := range mmp[tp[i]] {
			re = append(re, num)
			if len(re) == k {
				break LEE
			}
		}
	}
	return re
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}