package hard

import (
	"sort"
)

type MyCalendarThree struct {
	freq map[int]int
	st   []int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{make(map[int]int), []int{}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.record(start)
	this.record(end)

	sort.Ints(this.st)

	this.freq[start]++
	this.freq[end]--

	cnt, mx := 0, 0
	for _, v := range this.st {
		cnt += this.freq[v]
		if cnt > mx {
			mx = cnt
		}
	}

	return mx
}

func (this *MyCalendarThree) record(time int) {
	if _, ok := this.freq[time]; !ok {
		this.st = append(this.st, time)
	}
}
