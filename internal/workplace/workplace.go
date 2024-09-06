package workplace

import (
	"errors"
	"fmt"
)

var (
	ErrStartTimeIsNotIn15MinutesIncrements = errors.New("start time is not in 15 min increments")
	ErrEndTimeIsNotIn15MinutesIncrements   = errors.New("end time is not in 15 min increments")
	ErrPeriodOverlapsWithBufferTime        = errors.New("period overlaps with buffer time")
)

type Workplace struct {
	bufTimes []Period
	rooms    rooms
}

func New() *Workplace {
	return &Workplace{}
}

func (wp *Workplace) AddRoom(name string, capacity NumOfPeople) error {
	return addRoom(&wp.rooms, name, capacity)
}

func (wp *Workplace) AddBufferTime(p Period) {
	wp.bufTimes = append(wp.bufTimes, p)
}

func (wp *Workplace) RoomsAvailable(p Period) []Vacancy {
	if isInBufferTime(wp, p) {
		return nil
	}

	return findVacancies(wp.rooms, p)
}

func (wp *Workplace) Book(p Period, n NumOfPeople) (r Reservation, err error) {
	err = validateBooking(wp, p, n)
	if err != nil {
		return
	}

	r, err = findAndReserveRoom(wp.rooms, p, n)

	return
}

func validateBooking(wp *Workplace, p Period, n NumOfPeople) error {
	err := validatePeriod(wp, p)
	if err != nil {
		return &RoomReserveError{p, n, err}
	}
	return nil
}

func isInBufferTime(wp *Workplace, p Period) bool {
	return isAnyOverlapping(wp.bufTimes, p)
}

type PeriodValidationError struct {
	Period Period
	Err    error
}

func (err *PeriodValidationError) Error() string {
	return fmt.Sprintf("invalid period `%v`: %s", err.Period, err.Err)
}

func (err *PeriodValidationError) Unwrap() error {
	return err.Err
}

func validatePeriod(wp *Workplace, p Period) error {
	var errs []error

	if p.start.mm%15 != 0 {
		errs = append(errs, ErrStartTimeIsNotIn15MinutesIncrements)
	}

	if p.end.mm%15 != 0 {
		errs = append(errs, ErrEndTimeIsNotIn15MinutesIncrements)
	}

	if isInBufferTime(wp, p) {
		errs = append(errs, ErrPeriodOverlapsWithBufferTime)
	}

	if len(errs) > 0 {
		return &PeriodValidationError{p, errors.Join(errs...)}
	}

	return nil
}
