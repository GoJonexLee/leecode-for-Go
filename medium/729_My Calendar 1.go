package medium

type MyCalendar struct {
	calendar map[int]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		calendar: make(map[int]int),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for kStart, vEnd := range this.calendar {
		if !(kStart >= end || vEnd <= start) {
			return false
		}
	}
	this.calendar[start] = end
	return true
}
