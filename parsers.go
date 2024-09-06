package main

import (
	"geektrust/internal/workplace"
	"strconv"
)

func parseTime(s string) (time workplace.Time, err error) {
	var hh, mm int

	hh, err = strconv.Atoi(s[:2])
	if err != nil {
		return
	}

	mm, err = strconv.Atoi(s[3:])
	if err != nil {
		return
	}

	time, err = workplace.NewTime(uint8(hh), uint8(mm))

	return
}

func parsePeriod(start, end string) (period workplace.Period, err error) {
	var startTime, endTime workplace.Time

	startTime, err = parseTime(start)
	if err != nil {
		return
	}

	endTime, err = parseTime(end)
	if err != nil {
		return
	}

	period, err = workplace.NewPeriod(startTime, endTime)
	if err != nil {
		return
	}

	return
}
