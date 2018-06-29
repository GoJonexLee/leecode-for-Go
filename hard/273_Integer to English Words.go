package hard

import (
	"bytes"
)

var (
	v1 = [20]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	v2 = [10]string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	v3 = [3]string{"Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {

	dt := [][]byte{numToString(num % 1000).Bytes()}
	for i := 0; i < 3; i++ {
		num /= 1000
		if n := num % 1000; n != 0 {
			tmp := numToString(n)
			tmp.WriteString(" ")
			tmp.WriteString(v3[i])
			dt = append(dt, tmp.Bytes())
		}
	}

	for i := 0; i < len(dt)/2; i++ {
		dt[i], dt[len(dt)-1-i] = dt[len(dt)-1-i], dt[i]
	}

	re := bytes.Join(dt, []byte(" "))

	if len(re) == 0 {
		return "Zero"
	}
	return string(bytes.TrimSpace(re))
}

func numToString(num int) *bytes.Buffer {
	var (
		a = num / 100
		b = num % 100
		c = num % 10
	)

	var re = bytes.NewBufferString("")

	if a > 0 {
		re.WriteString(v1[a])
		re.WriteString(" Hundred")
	}

	if a > 0 && b > 0 {
		re.WriteString(" ")
	}

	if b < 20 {
		re.WriteString(v1[b])
	} else {
		re.WriteString(v2[b/10])
		if c > 0 {
			re.WriteString(" ")
			re.WriteString(v1[c])
		}
	}

	return re
}
