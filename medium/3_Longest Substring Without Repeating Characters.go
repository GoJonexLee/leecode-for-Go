package medium

func lengthOfLongestSubstring(s string) int {

	tmp := [256]int{}			// 记录上一次出现的位置
	for i := 0; i < 256; i++ {
		tmp[i] = -1
	}

	index, max := -1, 0
	for i := 0; i < len(s); i++ {
		//当前字符出现过，修改子串起始位置到该字符上一次出现的后一个位置（位置从-1开始）
		if tmp[s[i]] > index {
			index = tmp[s[i]]
		}
		//取最大值
		if i-index > max {
			max = i - index
		}
		//记录字符最近一次出现的位置
		tmp[s[i]] = i
	}
	return max
}