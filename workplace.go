package main

import "geektrust/internal/workplace"

func makeWorkplace() *workplace.Workplace {
	wp := workplace.New()

	bt1, _ := parsePeriod("09:00", "09:15")
	wp.AddBufferTime(bt1)

	bt2, _ := parsePeriod("13:15", "13:45")
	wp.AddBufferTime(bt2)

	bt3, _ := parsePeriod("18:45", "19:00")
	wp.AddBufferTime(bt3)

	_ = wp.AddRoom("C-Cave", 3)
	_ = wp.AddRoom("D-Tower", 7)
	_ = wp.AddRoom("G-Mansion", 20)

	return wp
}
