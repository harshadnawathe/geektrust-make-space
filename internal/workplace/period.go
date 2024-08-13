package workplace

import "fmt"

type Period struct {
	start, end Time
}

func NewPeriod(start Time, end Time) Period {
	return Period{start, end}
}

func isOverlapping(p1 Period, p2 Period) bool {
	return isTimeBefore(p1.start, p2.end) && isTimeBefore(p2.start, p1.end)
}

func isAnyOverlapping(periods []Period, p Period) bool {
	for _, period := range periods {
		if isOverlapping(period, p) {
			return true
		}
	}
	return false
}

func (p Period) String() string {
	return fmt.Sprintf("%s - %s", p.start, p.end)
}

type Time struct {
	hh, mm uint8
}

func NewTime(hh uint8, mm uint8) Time {
	return Time{hh, mm}
}

func isTimeBefore(t1, t2 Time) bool {
	return t1.hh < t2.hh || (t1.hh == t2.hh && t1.mm < t2.mm)
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.hh, t.mm)
}
