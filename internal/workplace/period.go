package workplace

import (
	"errors"
	"fmt"
)

type Period struct {
	start, end Time
}

func NewPeriod(start Time, end Time) (p Period, err error) {
	if isTimeBefore(end, start) {
		err = errors.New("end is before start")
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
	if hh > 23 {
		err = &TimeError{hh, mm, ErrTimeInvalidHourValue}
		return
	}

	if mm > 59 {
		err = &TimeError{hh, mm, ErrTimeInvalidMinuteValue}
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

type TimeValueError string

func (err TimeValueError) Error() string {
	return string(err)
}

const (
	ErrTimeInvalidHourValue   TimeValueError = "hour value must be between 0 and 23"
	ErrTimeInvalidMinuteValue TimeValueError = "minute value must be between 0 and 59"
)
