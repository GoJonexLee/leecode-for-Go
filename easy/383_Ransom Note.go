package easy

func canConstruct(ransomNote string, magazine string) bool {
	re := map[rune]int{}
	for _, m := range magazine {
		re[m]++
	}
	for _, r := range ransomNote {
		if num, ok := re[r]; !ok || num == 0 {
			return false
		} else {
			num--
			re[r] = num
		}
	}
	return true
}
