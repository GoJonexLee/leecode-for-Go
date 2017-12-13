package medium

import (
	"container/heap"
)

type Nums []int

func (n Nums) Len() int {
	return len(n)
}

func (n Nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nums) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n *Nums) Pop() interface{} {
	old := *n
	nc := len(old)
	x := old[nc-1]
	*n = old[:nc-1]
	return x
}

func (n *Nums) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n Nums) Top() int {
	return n[0]
}

func (n Nums) Set(num int) {
	n[0] = num
}

func findKthLargest(nums []int, k int) int {
	hp := &Nums{}

	for i := 0; i < len(nums); i++ {
		if hp.Len() == k {
			if nums[i] > hp.Top() {
                hp.Set(nums[i])
				heap.Fix(hp, 0)
			}
		} else {
			heap.Push(hp, nums[i])
		}
	}
	return hp.Top()
}