package easy

import "strings"

func reverseWords(s string) string {
	items := strings.Split(s, " ")
	re := []string{}

	for _, item := range items {
		re = append(re, reverse(item))
	}

	return strings.Join(re, " ")
}

func reverse(s string) string {
	st, lg := []byte(s), len(s)
	for i := 0; i < len(st)/2; i++ {
		st[i], st[lg-1-i] = st[lg-1-i], st[i]
	}
	return string(st)
}