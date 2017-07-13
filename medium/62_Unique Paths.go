package medium

func initArray(n, value int) []int {
	re := make([]int, n)

	for i := 0; i < n; i++ {
		re[i] = value
	}

	return re
}

func uniquePaths(m int, n int) int {

	if m == 1 || n == 1 {
		return 1
	}

	re := initArray(n, 1)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			re[j] = re[j-1] + re[j]
		}
	}

	return re[n-1]
}