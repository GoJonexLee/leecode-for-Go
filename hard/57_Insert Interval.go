package hard

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {

	re := []Interval{}
	if len(intervals) == 0 {
		re = append(re, newInterval)
		return re
	}

	var (
		idx = 0
		gap = Interval{-1, -1}
	)

	for i, dt := range intervals {
		if newInterval.Start > dt.End {
			re = append(re, dt)
			continue
		}

		if gap.Start != -1 {
			goto GEND
		}
		if newInterval.Start >= dt.Start && newInterval.Start <= dt.End {
			gap.Start = dt.Start
		} else if gap.Start == -1 && newInterval.Start < dt.Start {
			gap.Start = newInterval.Start
		}

		if newInterval.End > dt.End {
			continue
		}
	GEND:

		if newInterval.End > dt.End {
			continue
		} else if newInterval.End <= dt.End && newInterval.End >= dt.Start {
			gap.End = dt.End
			idx = i + 1
		} else if gap.End == -1 && newInterval.End < dt.Start {
			gap.End = newInterval.End
			idx = i
		}
		break
	}

	if gap.Start == -1 {
		gap.Start = newInterval.Start
	}

	if gap.End == -1 {
		gap.End = newInterval.End
		idx = len(intervals)
	}

	re = append(re, gap)
	re = append(re, intervals[idx:]...)
	return re
}
