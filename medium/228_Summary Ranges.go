package medium

import (
	"bytes"
	"strconv"
	"fmt"
)

func main() {
	s1 := []int{0,1,2,4,5,7}		// ["0->2","4->5","7"]
	s2 := []int{0,2,3,4,6,8,9}		// ["0","2->4","6","8->9"]

	fmt.Println(summaryRanges(s1))
	fmt.Println(summaryRanges(s2))
	fmt.Println(summaryRanges([]int{}))
}

func summaryRanges(nums []int) []string {

	re := []string{}

	for begin, end := 0, 0; end < len(nums); end++ {
		if end+1 < len(nums) && nums[end] + 1 == nums[end+1] {
			continue
		}

		if begin == end {
			re = append(re, strconv.Itoa(nums[begin]))
		} else {
			tmp := bytes.Buffer{}
			tmp.WriteString(strconv.Itoa(nums[begin]))
			tmp.WriteString("->")
			tmp.WriteString(strconv.Itoa(nums[end]))
			re = append(re, tmp.String())
		}
		begin = end + 1
	}

	return re
}