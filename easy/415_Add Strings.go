package easy

import (
	"container/list"
	"strings"
)

var mp = map[byte]int{'0':0, '1':1, '2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9}
var pm = map[int]string{0:"0", 1:"1", 2:"2", 3:"3", 4:"4", 5:"5", 6:"6", 7:"7", 8:"8", 9:"9"}


func addStrings(num1 string, num2 string) string {
	re, add := list.New(), 0
	i, j := len(num1)-1, len(num2)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		sum := mp[num1[i]] + mp[num2[j]] + add
		add = sum / 10
		re.PushFront(sum % 10)
	}

	for ; i >= 0; i-- {
		sum := mp[num1[i]] + add
		add = sum / 10
		re.PushFront(sum % 10)
	}

	for ; j >= 0; j-- {
		sum := mp[num2[j]] + add
		add = sum / 10
		re.PushFront(sum % 10)
	}

	if add > 0 {
		re.PushFront(add)
	}

	rest := []string{}
	for e := re.Front(); e != nil; e = e.Next() {
		if d, ok := e.Value.(int); ok {
			rest = append(rest, pm[d])
		}
	}
	return strings.Join(rest, "")
}
