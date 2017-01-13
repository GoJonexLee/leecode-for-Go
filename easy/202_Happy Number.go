package easy

import "strconv"

func isHappy(n int) bool {
	var f func(map[int]bool, int) bool
	f = func(mp map[int]bool, n int) bool {
		ns, tm := strconv.Itoa(n), 0

		for i := 0; i < len(ns); i++ {
			tn, _ := strconv.Atoi(string(ns[i]))
			tm += tn * tn
		}

		if tm == 1 {
			return true
		}

		if _, ok := mp[tm]; ok {
			return false
		}
		mp[tm] = true
		return f(mp, tm)
	}
	mpp := make(map[int]bool)
	return f(mpp, n)
}
