package easy

func firstUniqChar(s string) int {
	tmp := make(map[rune]int)
	mp := make(map[rune]int)
	for i, t := range s {
		tmp[t]++
		mp[t] = i
	}

	tp := make(map[int][]rune)
	for k, v := range tmp {
		tp[v] = append(tp[v], k)
	}

	if ls, ok := tp[1]; !ok {
		return -1
	} else {
		re := len(s)
		for _, l := range ls {
			if mp[l] < re {
				re = mp[l]
			}
		}
		return re
	}
}
