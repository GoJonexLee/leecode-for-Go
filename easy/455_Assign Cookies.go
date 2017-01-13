package easy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	k := 0
	for i, j := 0, 0; i < len(g) && j < len(s); i, j = i+1, j+1 {
		if g[i] <= s[j] {
			k++
		} else {
			i--
		}
	}
	return k
}