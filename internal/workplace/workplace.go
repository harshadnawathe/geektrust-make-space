package workplace

type Vacancy struct{ Room string }

type Period struct {
	start, end Time
}

type Workplace struct {
	bufTime []Period
	v       []Vacancy
}

func New() *Workplace {
	return &Workplace{}
}

func (wp *Workplace) AddRoom(s string) {
	wp.v = append(wp.v, Vacancy{Room: s})
}

func (wp *Workplace) AddBufferTime(p Period) {
	wp.bufTime = append(wp.bufTime, p)
}

func (wp *Workplace) AvailableRooms(p Period) []Vacancy {
	if len(wp.bufTime) > 0 {
		if isOverlapping(wp.bufTime[0], p) {
			return nil
		}
	}

	return wp.v
}

func NewPeriod(start Time, end Time) Period {
	return Period{start, end}
}

func isOverlapping(p1 Period, p2 Period) bool {
	return isTimeBefore(p1.start, p2.end) && isTimeBefore(p2.start, p1.end)
}

type Time struct {
	hh, mm uint8
}

func isTimeBefore(t1, t2 Time) bool {
	return t1.hh < t2.hh || (t1.hh == t2.hh && t1.mm < t2.mm)
}

func NewTime(hh uint8, mm uint8) Time {
	return Time{hh, mm}
}
