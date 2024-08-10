package workplace_test

import (
	"geektrust/internal/workplace"
	"reflect"
	"testing"
)

func Test_Workplace_AvailableRooms_Returns_EmptyList(t *testing.T) {
	w := workplace.New()

	vacancies := w.AvailableRooms(workplace.Period{})

	want := 0
	if got := len(vacancies); want != got {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_Returns_AllRooms(t *testing.T) {
	w := workplace.New()
	w.AddRoom("C-Cave")
	w.AddRoom("D-Tower")

	got := w.AvailableRooms(workplace.Period{})

	want := []workplace.Vacancy{{Room: "C-Cave"}, {Room: "D-Tower"}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_DuringBufferTime(t *testing.T) {
	w := workplace.New()
	w.AddRoom("C-Cave")
	w.AddRoom("D-Tower")
	w.AddBufferTime(workplace.NewPeriod(
		workplace.NewTime(9, 0), workplace.NewTime(9, 15),
	))
	w.AddBufferTime(workplace.NewPeriod(
		workplace.NewTime(13, 15), workplace.NewTime(13, 45),
	))

	tests := []struct {
		Name   string
		Period workplace.Period
		Want   int
	}{
		{
			Name:   "empty when start time is buffer time start",
			Period: workplace.NewPeriod(workplace.NewTime(9, 0), workplace.NewTime(10, 0)),
			Want:   0,
		},
		{
			Name:   "empty when end time is buffer time end",
			Period: workplace.NewPeriod(workplace.NewTime(8, 0), workplace.NewTime(9, 15)),
			Want:   0,
		},
		{
			Name:   "empty when start time in buffer time",
			Period: workplace.NewPeriod(workplace.NewTime(9, 5), workplace.NewTime(9, 30)),
			Want:   0,
		},
		{
			Name:   "empty when end time in buffer time",
			Period: workplace.NewPeriod(workplace.NewTime(8, 0), workplace.NewTime(9, 5)),
			Want:   0,
		},
		{
			Name:   "empty when buffer time in period",
			Period: workplace.NewPeriod(workplace.NewTime(13, 0), workplace.NewTime(14, 0)),
			Want:   0,
		},
		{
			Name:   "not empty when start time is buffer end time",
			Period: workplace.NewPeriod(workplace.NewTime(9, 15), workplace.NewTime(9, 30)),
			Want:   2,
		},
		{
			Name:   "not empty when end time is buffer start time",
			Period: workplace.NewPeriod(workplace.NewTime(8, 0), workplace.NewTime(9, 00)),
			Want:   2,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			vacancies := w.AvailableRooms(test.Period)

			if got := len(vacancies); test.Want != got {
				t.Errorf("AvailableRooms()= %v, want= %v", got, test.Want)
			}
		})
	}
}
