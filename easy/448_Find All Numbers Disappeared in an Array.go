package easy

/*

 */

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			nums[k] = i + 1
			k++
		}
	}
	return nums[:k]
}