package medium

import (
	"sort"
)

type MyCalendarTwo struct {
	time []int
	freq map[int]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		time: []int{},
		freq: make(map[int]int),
	}
}

func (this *MyCalendarTwo) removeTime(date int) {
	if freq := this.freq[date]; freq == 0 {
		delete(this.freq, date)
		var idx int
		for i, v := range this.time {
			if date == v {
				idx = i
				break
			}
		}
		this.time = append(this.time[0:idx], this.time[idx+1:]...)
	}
}

func (this *MyCalendarTwo) addTime(date int) {
	if _, ok := this.freq[date]; !ok {
		this.time = append(this.time, date)
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	this.addTime(start)
	this.addTime(end)

	if !sort.IntsAreSorted(this.time) {
		sort.Ints(this.time)
	}

	sFreq, eFreq := this.freq[start], this.freq[end]
	this.freq[start], this.freq[end] = sFreq+1, eFreq-1

	cnt := 0
	for _, v := range this.time {
		cnt += this.freq[v]
		if cnt >= 3 {
			this.freq[start], this.freq[end] = sFreq, eFreq
			this.removeTime(start)
			this.removeTime(end)
			return false
		}
	}

	return true
}
