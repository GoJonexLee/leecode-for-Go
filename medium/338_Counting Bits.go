package medium

func countBits(num int) []int {
	re := make([]int, num+1)

	for i := 1; i <= num; i++ {
		re[i] = re[i & (i-1)] + 1
	}

	return re
}
