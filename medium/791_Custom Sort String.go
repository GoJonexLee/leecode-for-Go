package medium

import (
	"bytes"
)

func customSortString(S string, T string) string {
	ct := make(map[rune]int)
	for _, t := range T {
		ct[t]++
	}

	re := bytes.NewBufferString("")
	for _, s := range S {
		for i := 0; i < ct[s]; i++ {
			re.WriteRune(s)
		}
		delete(ct, s)
	}

	for k, v := range ct {
		for i := 0; i < v; i++ {
			re.WriteRune(k)
		}
	}

	return re.String()
}
