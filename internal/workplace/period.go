package workplace

import (
	"errors"
	"fmt"
)

var (
	ErrTimeInvalidHourValue        = errors.New("hour value must be between 0 and 23")
	ErrTimeInvalidMinuteValue      = errors.New("minute value must be between 0 and 59")
	ErrPeriodValueEndIsBeforeStart = errors.New("end is before start")
)

type Period struct {
	start, end Time
}

func NewPeriod(start Time, end Time) (p Period, err error) {
	if isTimeBefore(end, start) {
		err = &PeriodError{start, end, ErrPeriodValueEndIsBeforeStart}
		return
	}

	p = Period{start, end}

	return
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

func NewTime(hh uint8, mm uint8) (t Time, err error) {
	var errs []error

	if hh > 23 {
		errs = append(errs, ErrTimeInvalidHourValue)
	}

	if mm > 59 {
		errs = append(errs, ErrTimeInvalidMinuteValue)
	}

	if errs != nil {
		err = &TimeError{hh, mm, errors.Join(errs...)}
		return
	}

	t = Time{hh, mm}

	return
}

func isTimeBefore(t1, t2 Time) bool {
	return t1.hh < t2.hh || (t1.hh == t2.hh && t1.mm < t2.mm)
}

func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.hh, t.mm)
}

type TimeError struct {
	HH, MM uint8
	Err    error
}

func (err *TimeError) Error() string {
	return fmt.Sprintf("invalid time value `%02d:%02d`: %s", err.HH, err.MM, err.Err)
}

func (err *TimeError) Unwrap() error {
	return err.Err
}

type PeriodError struct {
	Start, End Time
	Err        error
}

func (err *PeriodError) Error() string {
	return fmt.Sprintf("invalid period value `%s - %s`: %s", err.Start, err.End, err.Err)
}

func (err *PeriodError) Unwrap() error {
	return err.Err
}
