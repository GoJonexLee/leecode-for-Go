package easy

import (
	"sort"
)

func findLHS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

    count := make(map[int]int)
    for _, num:= range nums{
        count[num]++
    }
    
    if len(count) == 0 {
        return 0
	}
	
	var re int
    for k, v := range count {
        if tmp, ok :=count[k+1]; ok {
            if re < v + tmp {
                re = v + tmp
            }
        }
    }
    return re
}