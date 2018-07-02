package medium

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	re := NumMatrix{make([][]int, len(matrix)+1)}
	for i := 0; i < len(re.dp); i++ {
		re.dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i < len(re.dp); i++ {
		for j := 1; j < len(re.dp[i]); j++ {
			re.dp[i][j] = re.dp[i][j-1] + re.dp[i-1][j] - re.dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return re
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1] + this.dp[row1][col1]
}
