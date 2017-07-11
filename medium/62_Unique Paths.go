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

	cul := make(map[int][]int, 2)
	cul[0], cul[1] = initArray(n, 1), initArray(n, 1)

	var step, up int
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			step, up = i % 2, (i + 1) % 2
			cul[step][j] = cul[step][j-1] + cul[up][j]
		}
	}

	return cul[step][n-1]
}