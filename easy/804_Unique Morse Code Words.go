package easy

import (
	"bytes"
)

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	re := map[string]int

	for _, word := range words {
		tmp := bytes.NewBufferString("")
		for _, w := range word {
			tmp.WriteString(morse[w - 97])
		}
		re[tmp.String()]++
	}

	return len(re)
}
