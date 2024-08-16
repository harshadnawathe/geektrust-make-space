package workplace

import (
	"errors"
	"testing"
)

func TestPeriodValidationError_Error(t *testing.T) {
	err := &PeriodValidationError{
		Period: PeriodForTest("10:13", "10:24"),
		Err:    errors.New("some error"),
	}

	got := err.Error()

	want := "invalid period `10:13 - 10:24`: some error"
	if got != want {
		t.Errorf("PeriodValidationError.Error() = %v, want %v", got, want)
	}
}

func TestPeriodValidationError_Unwrap(t *testing.T) {
	err := &PeriodValidationError{Err: ErrStartTimeIsNotIn15MinutesIncrements}

	got := err.Unwrap()

	want := ErrStartTimeIsNotIn15MinutesIncrements

	if got != want {
		t.Errorf("PeriodValidationError.Unwrap() = %v, want %v", got, want)
	}
}
