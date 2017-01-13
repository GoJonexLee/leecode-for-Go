package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tmp := make(map[rune]int)
	for _, st := range s {
		tmp[st] += 1
	}

	for _, tt := range t {
		if v, ok := tmp[tt]; !ok || v <= 0 {
			return false
		}
		tmp[tt] -= 1
	}
	return true
}
