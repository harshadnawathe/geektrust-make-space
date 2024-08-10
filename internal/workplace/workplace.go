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
		if isTimeEqual(wp.bufTime[0].start, p.start) {
			return nil
		}
	}

	return wp.v
}

func NewPeriod(start Time, end Time) Period {
	return Period{start, end}
}

type Time struct {
	hh, mm uint8
}

func isTimeEqual(t1, t2 Time) bool {
	return t1.hh == t2.hh && t1.mm == t2.mm
}

func NewTime(hh uint8, mm uint8) Time {
	return Time{hh, mm}
}
