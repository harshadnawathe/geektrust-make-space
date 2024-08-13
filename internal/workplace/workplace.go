package workplace

import (
	"errors"
	"fmt"
	"sort"
)

type Workplace struct {
	bufTimes []Period
	rooms    []*room
}

func New() *Workplace {
	return &Workplace{}
}

func (wp *Workplace) AddRoom(name string, capacity NumOfPeople) {
	wp.rooms = append(wp.rooms, newRoom(name, capacity))

	sort.Slice(wp.rooms, func(i, j int) bool {
		return wp.rooms[i].capacity < wp.rooms[j].capacity
	})
}

func (wp *Workplace) AddBufferTime(p Period) {
	wp.bufTimes = append(wp.bufTimes, p)
}

func (wp *Workplace) AvailableRooms(p Period) []Vacancy {
	if isInBufferTime(wp, p) {
		return nil
	}

	return findVacancies(wp.rooms, p)
}

func (wp *Workplace) Book(p Period, n NumOfPeople) (r Reservation, err error) {
	err = validatePeriod(wp, p)
	if err != nil {
		err = fmt.Errorf("cannot book: %w", err)
		return
	}

	r, err = findAndReserveRoom(wp.rooms, p, n)
	if err != nil {
		err = fmt.Errorf("cannot book: %w", err)
		return
	}

	return
}

func isInBufferTime(wp *Workplace, p Period) bool {
	return isAnyOverlapping(wp.bufTimes, p)
}

func validatePeriod(wp *Workplace, p Period) error {
	if p.start.mm%15 != 0 {
		return errors.New("start time is not in 15 min increments")
	}

	if p.end.mm%15 != 0 {
		return errors.New("end time is not in 15 min increments")
	}

	if isInBufferTime(wp, p) {
		return errors.New("period overlaps with buffer time")
	}

	return nil
}
