package workplace

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	ErrRoomNameIsBlank        = errors.New("room name is blank")
	ErrRoomCapacityIsZero     = errors.New("room capacity is zero")
	ErrRoomAlreadyBooked      = errors.New("room is already booked")
	ErrRoomCapacityIsTooSmall = errors.New("room capacity is too small")
	ErrRoomNoVacantRoom       = errors.New("no vacant room")
)

type NumOfPeople uint

type room struct {
	name     string
	bookedAt []Period
	capacity NumOfPeople
}

type rooms []*room

type RoomInitError struct {
	Name     string
	Capacity NumOfPeople
	Err      error
}

func (err *RoomInitError) Error() string {
	return fmt.Sprintf(
		"cannot create a room with name `%v` and capacity `%v`: %s",
		err.Name,
		err.Capacity,
		err.Err,
	)
}

func (err *RoomInitError) Unwrap() error {
	return err.Err
}

func newRoom(name string, capacity NumOfPeople) (*room, error) {
	var errs []error

	if len(strings.TrimSpace(name)) == 0 {
		errs = append(errs, ErrRoomNameIsBlank)
	}

	if capacity == 0 {
		errs = append(errs, ErrRoomCapacityIsZero)
	}

	if len(errs) > 0 {
		return nil, &RoomInitError{
			Name:     name,
			Capacity: capacity,
			Err:      errors.Join(errs...),
		}
	}

	return &room{
		name:     name,
		capacity: capacity,
	}, nil
}

func isBooked(r *room, p Period) bool {
	return isAnyOverlapping(r.bookedAt, p)
}

type Reservation struct {
	Room string
}

type RoomReserveError struct {
	Period      Period
	NumOfPeople NumOfPeople
	Err         error
}

func (err *RoomReserveError) Error() string {
	return fmt.Sprintf(
		"cannot reserve room for `%d` people in period `%v`: %s",
		err.NumOfPeople,
		err.Period,
		err.Err,
	)
}

func (err *RoomReserveError) Unwrap() error {
	return err.Err
}

func reserve(r *room, p Period, n NumOfPeople) (res Reservation, err error) {
	err = validateCapacity(r, n)
	if err != nil {
		err = &RoomReserveError{
			Period:      p,
			NumOfPeople: n,
			Err:         err,
		}
		return
	}

	if isBooked(r, p) {
		err = &RoomReserveError{
			Period:      p,
			NumOfPeople: n,
			Err:         ErrRoomAlreadyBooked,
		}
		return
	}

	r.bookedAt = append(r.bookedAt, p)

	res = Reservation{r.name}

	return
}

type RoomCapacityValidationError struct {
	Name        string
	Capacity    NumOfPeople
	NumOfPeople NumOfPeople
	Err         error
}

func (err *RoomCapacityValidationError) Error() string {
	return fmt.Sprintf(
		"cannot fit `%d` people in room `%s` with capacity `%d`: %s",
		err.NumOfPeople,
		err.Name,
		err.Capacity,
		err.Err,
	)
}

func (err *RoomCapacityValidationError) Unwrap() error {
	return err.Err
}

func validateCapacity(r *room, n NumOfPeople) error {
	if n > r.capacity {
		return &RoomCapacityValidationError{r.name, r.capacity, n, ErrRoomCapacityIsTooSmall}
	}

	return nil
}

func addRoom(rooms *rooms, name string, capacity NumOfPeople) error {
	r, _ := newRoom(name, capacity)

	*rooms = append(*rooms, r)

	sort.Slice(*rooms, func(i, j int) bool {
		return (*rooms)[i].capacity < (*rooms)[j].capacity
	})

	return nil
}

func findAndReserveRoom(rooms rooms, p Period, n NumOfPeople) (Reservation, error) {
	for _, room := range rooms {
		if reservation, err := reserve(room, p, n); err == nil {
			return reservation, nil
		}
	}

	return Reservation{}, &RoomReserveError{p, n, ErrRoomNoVacantRoom}
}

type Vacancy struct {
	Room string
}

func findVacancies(rooms rooms, p Period) []Vacancy {
	var vacancies []Vacancy
	for _, room := range rooms {
		if !isBooked(room, p) {
			vacancies = append(vacancies, Vacancy{room.name})
		}
	}
	return vacancies
}
