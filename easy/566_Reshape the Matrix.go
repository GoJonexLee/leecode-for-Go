package easy

func matrixReshape(nums [][]int, r int, c int) [][]int {

	if len(nums) * len(nums[0]) != r * c {
		return nums
	}

	re, row, count := make([][]int, r), 0, 0
	for _, num := range nums {
		for _, nu := range num {
			if count == 0 {
				re[row] = []int{}
			} else if count == c {
				count = 0
				row++
			}
			count++
			re[row] = append(re[row], nu)
		}
	}

	return re
}