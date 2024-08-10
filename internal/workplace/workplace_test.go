package workplace_test

import (
	"geektrust/internal/workplace"
	"reflect"
	"testing"
)

func Test_Workplace_AvailableRooms_Returns_EmptyList(t *testing.T) {
	w := workplace.New()

	vacancies := w.AvailableRooms(workplace.PeriodForTest("10:00", "12:00"))

	want := 0
	if got := len(vacancies); want != got {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_Returns_AllRooms(t *testing.T) {
	w := workplace.New()
	w.AddRoom("C-Cave")
	w.AddRoom("D-Tower")

	got := w.AvailableRooms(workplace.PeriodForTest("10:00", "12:00"))

	want := []workplace.Vacancy{{Room: "C-Cave"}, {Room: "D-Tower"}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_DuringBufferTime(t *testing.T) {
	w := workplace.New()
	w.AddRoom("C-Cave")
	w.AddRoom("D-Tower")
	w.AddBufferTime(workplace.PeriodForTest("09:00", "09:15"))
	w.AddBufferTime(workplace.PeriodForTest("13:15", "13:45"))

	tests := []struct {
		Name   string
		Period workplace.Period
		Want   int
	}{
		{
			Name:   "empty when start time is buffer time start",
			Period: workplace.PeriodForTest("09:00", "10:00"),
			Want:   0,
		},
		{
			Name:   "empty when end time is buffer time end",
			Period: workplace.PeriodForTest("08:00", "09:15"),
			Want:   0,
		},
		{
			Name:   "empty when start time in buffer time",
			Period: workplace.PeriodForTest("09:05", "10:00"),
			Want:   0,
		},
		{
			Name:   "empty when end time in buffer time",
			Period: workplace.PeriodForTest("08:00", "09:05"),
			Want:   0,
		},
		{
			Name:   "empty when buffer time in period",
			Period: workplace.PeriodForTest("13:00", "14:00"),
			Want:   0,
		},
		{
			Name:   "not empty when start time is buffer end time",
			Period: workplace.PeriodForTest("09:15", "10:00"),
			Want:   2,
		},
		{
			Name:   "not empty when end time is buffer start time",
			Period: workplace.PeriodForTest("08:00", "09:00"),
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
