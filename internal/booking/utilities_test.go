package booking

import "geektrust/internal/workplace"

func NewPeriodMust(p workplace.Period, err error) workplace.Period {
	if err != nil {
		panic(err)
	}
	return p
}

func NewTimeMust(p workplace.Time, err error) workplace.Time {
	if err != nil {
		panic(err)
	}
	return p
}
