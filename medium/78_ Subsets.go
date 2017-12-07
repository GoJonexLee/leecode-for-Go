package medium

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var deep func([]int, int, *[][]int)
    deep = func(tmp []int, in int, result *[][]int) {
		tp := make([]int, len(tmp))
		copy(tp, tmp)
		*result = append(*result, tp)
        for i := in; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
            deep(tmp, i+1, result)
			tmp = tmp[:len(tmp)-1]
		}
    }
    
    tmp, re := []int{}, [][]int{}
    deep(tmp, 0, &re)
    return re
}