package medium

func findDuplicate(nums []int) int {
    
    slow, fast := 0, 0
    for {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            break
        }
	}

	fmt.Println(slow)
	
    var finder int
    for {
        slow = nums[slow]
        finder = nums[finder]
        
        if slow == finder {
            break
        }
    }
    
    return slow
}

func findDuplicate1(nums []int) int {
	
	var re, tmp int
	for i := 0; i < len(nums); i++ {
		tmp = nums[i]

		if tmp < 0 {
			tmp = -tmp
		}

		if nums[tmp-1] < 0 {
			re = tmp
		} else {
			nums[tmp-1] = -nums[tmp-1]
		}
		fmt.Println(nums)
	}
    
	return re
}