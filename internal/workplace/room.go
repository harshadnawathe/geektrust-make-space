package workplace

import (
	"errors"
	"fmt"
)

type NumOfPeople uint

type room struct {
	name     string
	bookedAt []Period
	capacity NumOfPeople
}

func newRoom(name string, capacity NumOfPeople) *room {
	return &room{
		name:     name,
		capacity: capacity,
	}
}

func canFit(r *room, n NumOfPeople) bool {
	return n <= r.capacity
}

func isBooked(r *room, p Period) bool {
	return isAnyOverlapping(r.bookedAt, p)
}

type Reservation struct {
	Room string
}

func reserve(r *room, p Period, n NumOfPeople) (res Reservation, err error) {
	if !canFit(r, n) {
		err = fmt.Errorf("cannot reserve: room with capacity %v cannot fit %v people", r.capacity, n)
		return
	}

	if isBooked(r, p) {
		err = errors.New("cannot reserve: room is booked")
		return
	}

	r.bookedAt = append(r.bookedAt, p)

	res = Reservation{r.name}

	return
}

func findAndReserveRoom(rooms []*room, p Period, n NumOfPeople) (Reservation, error) {
	for _, room := range rooms {
		if reservation, err := reserve(room, p, n); err == nil {
			return reservation, nil
		}
	}

	return Reservation{}, errors.New("no vacant room")
}

type Vacancy struct {
	Room string
}

func findVacancies(rooms []*room, p Period) []Vacancy {
	var vacancies []Vacancy
	for _, room := range rooms {
		if !isBooked(room, p) {
			vacancies = append(vacancies, Vacancy{room.name})
		}
	}
	return vacancies
}
