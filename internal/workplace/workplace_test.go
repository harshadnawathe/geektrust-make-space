package workplace_test

import (
	"geektrust/internal/workplace"
	"reflect"
	"testing"
)

func Test_Workplace_AvailableRooms_Returns_EmptyList(t *testing.T) {
	w := workplace.Build(
		workplace.Default(),
		workplace.WithNoRooms(),
	)

	vacancies := w.AvailableRooms(workplace.PeriodForTest("10:00", "12:00"))

	want := 0
	if got := len(vacancies); want != got {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_Returns_AllRooms(t *testing.T) {
	w := workplace.Build(
		workplace.WithRoom("C-Cave", 3),
		workplace.WithRoom("D-Tower", 7),
	)

	got := w.AvailableRooms(workplace.PeriodForTest("10:00", "12:00"))

	want := []workplace.Vacancy{{Room: "C-Cave"}, {Room: "D-Tower"}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_Returns_AllRooms_InSortedOrderOfCapacity(t *testing.T) {
	w := workplace.Build(
		workplace.WithRoom("D-Tower", 7),
		workplace.WithRoom("C-Cave", 3),
	)

	got := w.AvailableRooms(workplace.PeriodForTest("10:00", "12:00"))

	want := []workplace.Vacancy{{Room: "C-Cave"}, {Room: "D-Tower"}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("AvailableRooms()= %v, want= %v", got, want)
	}
}

func Test_Workplace_AvailableRooms_DuringBufferTime(t *testing.T) {
	w := workplace.Build(
		workplace.WithRoom("C-Cave", 3),
		workplace.WithRoom("D-Tower", 7),
		workplace.WithBufferTime("09:00", "09:15"),
		workplace.WithBufferTime("13:15", "13:45"),
	)

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

func Test_Workplace_Book_Returns_Reservation(t *testing.T) {
	w := workplace.Build(workplace.WithRoom("C-Cave", 3))

	got, _ := w.Book(workplace.PeriodForTest("10:00", "12:00"), 2)

	want := workplace.Reservation{"C-Cave"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Book()= %v, want= %v", got, want)
	}
}

func Test_Workplace_Book_Reserves_First_Available_Room_That_Can_Fit_Given_People(t *testing.T) {
	w := workplace.Build(
		workplace.WithRoom("C-Cave", 3),
		workplace.WithRoom("D-Tower", 7),
		workplace.WithRoom("G-Mansion", 20),
	)

	got, _ := w.Book(workplace.PeriodForTest("10:00", "12:00"), 5)

	want := workplace.Reservation{"D-Tower"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Book()= %v, want= %v", got, want)
	}
}

func Test_Workplace_Book_Reserves_Next_Bigger_Room_When_A_Smaller_Room_Is_Unavailable(t *testing.T) {
	w := workplace.Build(
		workplace.WithRoom("C-Cave", 3),
		workplace.WithRoom("D-Tower", 7),
	)
	_, _ = w.Book(workplace.PeriodForTest("10:00", "12:00"), 2)

	got, _ := w.Book(workplace.PeriodForTest("10:30", "11:30"), 2)

	want := workplace.Reservation{"D-Tower"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Book()= %v, want= %v", got, want)
	}
}

func Test_Workplace_Book_DuringBufferTime(t *testing.T) {
	tests := []struct {
		Name     string
		Period   workplace.Period
		IsErrNil bool
	}{
		{
			Name:     "err when start time is buffer time start",
			Period:   workplace.PeriodForTest("09:00", "10:00"),
			IsErrNil: false,
		},
		{
			Name:     "err when end time is buffer time end",
			Period:   workplace.PeriodForTest("08:00", "09:15"),
			IsErrNil: false,
		},
		{
			Name:     "err when start time in buffer time",
			Period:   workplace.PeriodForTest("09:05", "10:00"),
			IsErrNil: false,
		},
		{
			Name:     "err when end time in buffer time",
			Period:   workplace.PeriodForTest("08:00", "09:05"),
			IsErrNil: false,
		},
		{
			Name:     "err when buffer time in period",
			Period:   workplace.PeriodForTest("13:00", "14:00"),
			IsErrNil: false,
		},
		{
			Name:     "not err when start time is buffer end time",
			Period:   workplace.PeriodForTest("09:15", "10:00"),
			IsErrNil: true,
		},
		{
			Name:     "not err when end time is buffer start time",
			Period:   workplace.PeriodForTest("08:00", "09:00"),
			IsErrNil: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := workplace.Build(
				workplace.WithRoom("D-Tower", 7),
				workplace.WithBufferTime("09:00", "09:15"),
				workplace.WithBufferTime("13:15", "13:45"),
			)

			_, err := w.Book(test.Period, 2)

			if got := err == nil; test.IsErrNil != got {
				t.Errorf("Book()= %v, want= %v", got, test.IsErrNil)
			}
		})
	}
}

func Test_Workplace_AvailableRooms_After_Book(t *testing.T) {
	tests := []struct {
		Name   string
		Want   []workplace.Vacancy
		Period workplace.Period
	}{
		{
			Name:   "when start time is busy time start",
			Period: workplace.PeriodForTest("10:00", "10:30"),
			Want:   []workplace.Vacancy{{"D-Tower"}},
		},
		{
			Name:   "when end time is busy time end",
			Period: workplace.PeriodForTest("14:30", "15:00"),
			Want:   []workplace.Vacancy{{"C-Cave"}},
		},
		{
			Name:   "when start time in busy time",
			Period: workplace.PeriodForTest("10:30", "12:00"),
			Want:   []workplace.Vacancy{{"D-Tower"}},
		},
		{
			Name:   "when end time in busy time",
			Period: workplace.PeriodForTest("13:45", "14:30"),
			Want:   []workplace.Vacancy{{"C-Cave"}},
		},
		{
			Name:   "when busy time in period",
			Period: workplace.PeriodForTest("09:00", "12:00"),
			Want:   []workplace.Vacancy{{"D-Tower"}},
		},
		{
			Name:   "when start time is busy end time",
			Period: workplace.PeriodForTest("09:15", "10:00"),
			Want:   []workplace.Vacancy{{"C-Cave"}, {"D-Tower"}},
		},
		{
			Name:   "when end time is busy start time",
			Period: workplace.PeriodForTest("08:00", "09:00"),
			Want:   []workplace.Vacancy{{"C-Cave"}, {"D-Tower"}},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := workplace.Build(
				workplace.WithRoom("C-Cave", 3),
				workplace.WithRoom("D-Tower", 7),
			)
			_, _ = w.Book(workplace.PeriodForTest("10:00", "11:00"), 2)
			_, _ = w.Book(workplace.PeriodForTest("14:00", "15:00"), 5)

			got := w.AvailableRooms(test.Period)

			if !reflect.DeepEqual(got, test.Want) {
				t.Errorf("AvailableRooms()= %v, want= %v", got, test.Want)
			}
		})
	}
}

func Test_Workplace_Book_Returns_Error_When_No_Rooms_Are_Available(t *testing.T) {
	w := workplace.Build(workplace.WithNoRooms())

	_, err := w.Book(workplace.PeriodForTest("10:00", "11:00"), 2)

	wantErr := true
	if gotErr := err != nil; gotErr != wantErr {
		t.Errorf("Book()= %v, wantErr= %v", gotErr, wantErr)
	}
}

func Test_Workplace_Book_Checks_If_Period_In_15_Minutes_Increments(t *testing.T) {
	type args struct {
		p           workplace.Period
		numOfPeople int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error when start time is not in 15 minutes increments",
			args: args{
				p:           workplace.PeriodForTest("10:03", "11:00"),
				numOfPeople: 2,
			},
			wantErr: true,
		},
		{
			name: "no error when start time is not in 15 minutes increments",
			args: args{
				p:           workplace.PeriodForTest("10:00", "11:00"),
				numOfPeople: 2,
			},
			wantErr: false,
		},
		{
			name: "error when end time is not in 15 minutes increments",
			args: args{
				p:           workplace.PeriodForTest("10:00", "11:16"),
				numOfPeople: 2,
			},
			wantErr: true,
		},
		{
			name: "no error when end time is not in 15 minutes increments",
			args: args{
				p:           workplace.PeriodForTest("10:00", "11:00"),
				numOfPeople: 2,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := workplace.Build(workplace.Default())

			_, err := w.Book(test.args.p, test.args.numOfPeople)

			if test.wantErr != (err != nil) {
				t.Errorf("Book() error= %v, wantErr= %v", err, test.wantErr)
			}
		})
	}
}
