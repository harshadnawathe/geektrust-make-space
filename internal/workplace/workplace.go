package workplace

import (
	"errors"
	"fmt"
	"sort"
)

type Vacancy struct {
	Room string
}

type Reservation struct {
	Room string
}

type Workplace struct {
	bufTimes []Period
	rooms    []*room
}

func New() *Workplace {
	return &Workplace{}
}

func (wp *Workplace) AddRoom(name string, capacity int) {
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

	var vacancies []Vacancy
	for _, room := range wp.rooms {
		if !isBooked(room, p) {
			vacancies = append(vacancies, Vacancy{room.name})
		}
	}
	return vacancies
}

func (wp *Workplace) Book(p Period, numOfPeople int) (r Reservation, err error) {
	err = validatePeriod(wp, p)
	if err != nil {
		err = fmt.Errorf("cannot book: %w", err)
		return
	}

	var bookedRoom *room
	bookedRoom, err = findAndBookRoom(wp.rooms, p, numOfPeople)
	if err != nil {
		err = fmt.Errorf("cannot book: %w", err)
		return
	}

	r = Reservation{bookedRoom.name}

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

