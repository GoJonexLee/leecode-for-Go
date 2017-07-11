package medium

func minPathSum(grid [][]int) int {
    
	re, size := make(map[int][]int, 2), len(grid[0])
	re[0], re[1] = make([]int, size), make([]int, size)

	var step, up int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			step, up = i % 2, (i + 1) % 2
			if i == 0 {
				if j == 0 {
					re[step][j] = grid[i][j]
				} else {
					re[step][j] = re[step][j-1] + grid[i][j]
				}
			} else if j == 0 {
				re[step][j] = re[up][j] + grid[i][j]
			} else {
				re[step][j] = min(re[step][j-1], re[up][j]) + grid[i][j]
			}
		}
	}
	
	return re[step][size - 1]
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}