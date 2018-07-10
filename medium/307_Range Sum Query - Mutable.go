package medium

type NumArray struct {
	num []int
	bit []int
}

func Constructor(nums []int) NumArray {
	nm := NumArray{
		num: make([]int, len(nums)+1),
		bit: make([]int, len(nums)+1),
	}

	for i, num := range nums {
		nm.Update(i, num)
	}

	return nm
}

func (this *NumArray) Update(i int, val int) {
	nval := val - this.num[i+1]
	for j := i + 1; j < len(this.num); j += (j & -j) {
		this.bit[j] += nval
	}
	this.num[i+1] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum(j+1) - this.sum(i)
}

func (nm *NumArray) sum(i int) int {
	re := 0
	for j := i; j > 0; j -= (j & -j) {
		re += nm.bit[j]
	}
	return re
}
