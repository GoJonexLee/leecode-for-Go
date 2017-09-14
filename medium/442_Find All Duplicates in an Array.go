package medium

func findDuplicates(nums []int) []int {
    
    re := []int{}
    
    var tmp int
    for i := 0; i < len(nums); i++ {
        tmp = nums[i]
		
		if tmp < 0 {
			tmp = -tmp
		}

		if nums[tmp-1] < 0 {
			re = append(re, tmp)
		}
		nums[tmp-1] = -nums[tmp-1]
    }
    
    return re
}