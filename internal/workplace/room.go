package workplace

import (
	"errors"
	"fmt"
)

type room struct {
	name     string
	bookedAt []Period
	capacity int
}

func newRoom(name string, capacity int) *room {
	return &room{
		name:     name,
		capacity: capacity,
	}
}

func canFit(r *room, numOfPeople int) bool {
	return numOfPeople <= r.capacity
}

func isBooked(r *room, p Period) bool {
	return isAnyOverlapping(r.bookedAt, p)
}

type Reservation struct {
	Room string
}

func reserve(r *room, p Period, numOfPeople int) (res Reservation, err error) {
	if !canFit(r, numOfPeople) {
		err = fmt.Errorf("cannot reserve: room with capacity %v cannot fit %v people", r.capacity, numOfPeople)
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

func findAndReserveRoom(rooms []*room, p Period, numOfPeople int) (Reservation, error) {
	for _, room := range rooms {
		if reservation, err := reserve(room, p, numOfPeople); err == nil {
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
