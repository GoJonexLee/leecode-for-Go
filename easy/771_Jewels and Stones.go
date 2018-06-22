package easy

func numJewelsInStones(J string, S string) int {
	ct := make(map[rune]int)
	for _, s := range S {
		ct[s]++
	}

	re := 0
	for _, j := range J {
		re += ct[j]
	}

	return re
}
