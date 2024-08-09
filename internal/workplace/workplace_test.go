package workplace_test

import (
	"geektrust/internal/workplace"
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


