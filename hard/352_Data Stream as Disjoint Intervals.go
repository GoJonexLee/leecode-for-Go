package hard

import (
	"container/list"
)

type Interval struct {
	Start int
 	End   int
}

type SummaryRanges struct {
	dt list.List
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) Addnum(val int)  {
	var pos *list.Element
	for e := this.dt.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == val {
			return
		} else if e.Value.(int) > val {
			pos = e
			break
		}
	}

	if pos == nil {
		this.dt.PushBack(val)
	} else {
		this.dt.InsertBefore(val, pos)
	}
}


func (this *SummaryRanges) Getintervals() []Interval {
	re := make([]Interval, 0)

	var p, begin, end *list.Element
	for begin, p,  end= this.dt.Front(), this.dt.Front(), this.dt.Front().Next(); end != nil; end = end.Next(){
		if end.Value.(int) - p.Value.(int) > 1 {
			re = append(re, Interval{begin.Value.(int), p.Value.(int)})
			begin, p = end, end
		} else {
			p = p.Next()
		}
	}
	re = append(re, Interval{begin.Value.(int), p.Value.(int)})
	return re
}
