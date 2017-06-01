package easy

func islandPerimeter(grid [][]int) int {
	re, rows, cols := 0, len(grid), len(grid[0])

	fc := func(i, j int) int {
		r := 0
		if grid[i][j] == 1 {
			if i == 0 {
				r++
			}
			if i == rows-1 {
				r++
			}
			if j == 0 {
				r++
			}
			if j == cols-1 {
				r++
			}

			if i-1 >= 0 && grid[i-1][j] == 0 {
				r++
			}
			if i+1 < rows && grid[i+1][j] == 0 {
				r++
			}
			if j-1 >= 0 && grid[i][j-1] == 0 {
				r++
			}
			if j+1 < cols && grid[i][j+1] == 0 {
				r++
			}

		}
		return r
	}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			re += fc(i, j)
		}
	}

	return re
}