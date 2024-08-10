package workplace

import "sort"

type Vacancy struct {
	Room string
}

type Period struct {
	start, end Time
}

type Workplace struct {
	bufTimes []Period
	rooms    []room
}

func New() *Workplace {
	return &Workplace{}
}

func (wp *Workplace) AddRoom(name string, capacity int) {
	wp.rooms = append(wp.rooms, room{name, capacity})
	sort.Slice(wp.rooms, func(i, j int) bool {
		return wp.rooms[i].capacity < wp.rooms[j].capacity
	})
}

func (wp *Workplace) AddBufferTime(p Period) {
	wp.bufTimes = append(wp.bufTimes, p)
}

func (wp *Workplace) AvailableRooms(p Period) []Vacancy {
	for _, bufTime := range wp.bufTimes {
		if isOverlapping(bufTime, p) {
			return nil
		}
	}

	var vacancies []Vacancy
	for _, room := range wp.rooms {
		vacancies = append(vacancies, Vacancy{room.name})
	}
	return vacancies
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

type room struct {
	name     string
	capacity int
}
