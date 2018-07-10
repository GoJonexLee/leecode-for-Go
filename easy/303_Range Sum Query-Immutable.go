package easy

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	re := NumArray{make([]int, len(nums)+1)}
	for i := 1; i <= len(nums); i++ {
		re.dp[i] = re.dp[i-1] + nums[i-1]
	}
	return re
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j+1] - this.dp[i]
}
