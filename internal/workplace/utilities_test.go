package workplace

import (
	"fmt"
	"strconv"
)

func NewTimeMust(hh, mm uint8) Time {
	time, err := NewTime(hh, mm)
	if err != nil {
		panic(err)
	}

	return time
}

func PeriodForTest(start, end string) Period {
	return NewPeriod(TimeForTest(start), TimeForTest(end))
}

func TimeForTest(time string) Time {
	var hh, mm int
	var err error

	hh, err = strconv.Atoi(time[:2])
	if err != nil {
		panic(fmt.Errorf("cannot parse time `%s`: %w", time, err))
	}
	mm, err = strconv.Atoi(time[3:])
	if err != nil {
		panic(fmt.Errorf("cannot parse time `%s`: %w", time, err))
	}

	return NewTimeMust(uint8(hh), uint8(mm))
}

type workplaceConfigurer func(w *Workplace)

func Default() workplaceConfigurer {
	return func(w *Workplace) {
		w.AddBufferTime(PeriodForTest("09:00", "09:15"))
		w.AddBufferTime(PeriodForTest("13:15", "13:45"))
		w.AddBufferTime(PeriodForTest("18:45", "19:00"))

		w.AddRoom("C-Cave", 3)
		w.AddRoom("D-Tower", 7)
		w.AddRoom("G-Mansion", 20)
	}
}

func WithNoRooms() workplaceConfigurer {
	return func(w *Workplace) {
		w.rooms = nil
	}
}

func WithRoom(name string, capacity int) workplaceConfigurer {
	return func(w *Workplace) {
		w.AddRoom(name, capacity)
	}
}

func WithBufferTime(start, end string) workplaceConfigurer {
	return func(w *Workplace) {
		w.AddBufferTime(PeriodForTest(start, end))
	}
}

func Build(configs ...workplaceConfigurer) *Workplace {
	w := New()
	for _, config := range configs {
		config(w)
	}
	return w
}
