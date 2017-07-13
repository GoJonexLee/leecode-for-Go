package medium

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func minPathSum(grid [][]int) int {
    
	size := len(grid[0])
	re := make([]int, size)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j == 0 {
					re[j] = grid[i][j]
				} else {
					re[j] = re[j-1] + grid[i][j]
				}
			} else if j == 0 {
				re[j] = re[j] + grid[i][j]
			} else {
				re[j] = min(re[j-1], re[j]) + grid[i][j]
			}
		}
	}

	return re[size-1]
}