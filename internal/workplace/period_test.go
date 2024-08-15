package workplace

import (
	"errors"
	"reflect"
	"testing"
)

func TestTime_String(t *testing.T) {
	type fields struct {
		hh uint8
		mm uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "convert time value to string",
			fields: fields{12, 30},
			want:   "12:30",
		},
		{
			name:   "prefix with 0 if hour value is single digit",
			fields: fields{9, 30},
			want:   "09:30",
		},
		{
			name:   "prefix with 0 if minute value is single digit",
			fields: fields{12, 9},
			want:   "12:09",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Time{
				hh: tt.fields.hh,
				mm: tt.fields.mm,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("Time.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeriod_String(t *testing.T) {
	p := PeriodForTest("09:30", "13:45")

	got := p.String()

	want := "09:30 - 13:45"
	if got != want {
		t.Errorf("String()= %v, want= %v", got, want)
	}
}

func TestNewTime(t *testing.T) {
	type args struct {
		hh uint8
		mm uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Time
		wantErr error
	}{
		{
			name: "create time with no error",
			args: args{
				hh: 11,
				mm: 59,
			},
			want:    Time{11, 59},
			wantErr: nil,
		},
		{
			name: "error when hour value is greater than 23",
			args: args{
				hh: 24,
				mm: 59,
			},
			want:    Time{},
			wantErr: ErrTimeInvalidHourValue,
		},
		{
			name: "error when minute value is greater than 59",
			args: args{
				hh: 12,
				mm: 60,
			},
			want:    Time{},
			wantErr: ErrTimeInvalidMinuteValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTime(tt.args.hh, tt.args.mm)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTime() = %v, want %v", got, tt.want)
			}

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NewTime() error = %vi, wantErr %v", err, tt.wantErr)
				}

				var timeError *TimeError
				if !errors.As(err, &timeError) {
					t.Errorf("cannot use error %v as *TimeError", err)
				}
				
				if timeError.HH != tt.args.hh {
					t.Errorf("timeError.HH = %v, want = %v", timeError.HH, tt.args.hh)
				}

				if timeError.MM != tt.args.mm {
          t.Errorf("timeError.MM = %v, want = %v", timeError.MM, tt.args.mm)
        }
			}
		})
	}
}

func TestNewPeriod(t *testing.T) {
	type args struct {
		start Time
		end   Time
	}
	tests := []struct {
		name    string
		args    args
		want    Period
		wantErr error
	}{
		{
			name: "create period with no error",
			args: args{
				start: TimeForTest("09:30"),
				end:   TimeForTest("13:45"),
			},
			want:    Period{Time{9, 30}, Time{13, 45}},
			wantErr: nil,
		},
		{
			name: "error when end is before start",
			args: args{
				start: TimeForTest("14:00"),
				end:   TimeForTest("13:00"),
			},
			want: Period{},
			wantErr: ErrPeriodValueEndIsBeforeStart,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPeriod(tt.args.start, tt.args.end)
			
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPeriod() = %v, want %v", got, tt.want)
			}

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NewPeriod() error = %v, wantErr %v", err, tt.wantErr)
				}

				var periodError *PeriodError
        if !errors.As(err, &periodError) {
          t.Errorf("cannot use error %v as *PeriodError", err)
        }

				if periodError.Start != tt.args.start {
					t.Errorf("periodError.Start= %v, want= %v", periodError.Start, tt.args.start)
				}

				if periodError.End != tt.args.end {
					t.Errorf("periodError.End= %v, want= %v", periodError.End, tt.args.end)
				}
			}
		})
	}
}

func TestTimeError_Error(t *testing.T) {
	type fields struct {
		HH  uint8
		MM  uint8
		Err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"invalid hour value",
			fields{24, 59, ErrTimeInvalidHourValue},
			"invalid time value `24:59`: hour value must be between 0 and 23",
		},
		{
			"invalid minute value",
			fields{12, 60, ErrTimeInvalidMinuteValue},
			"invalid time value `12:60`: minute value must be between 0 and 59",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &TimeError{
				HH:  tt.fields.HH,
				MM:  tt.fields.MM,
				Err: tt.fields.Err,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("TimeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeError_Unwrap(t *testing.T) {
	err := &TimeError{Err: ErrTimeInvalidHourValue}

	got := err.Unwrap()

	want := ErrTimeInvalidHourValue
	if got != want {
		t.Errorf("Unwrap()= %v, want= %v", got, want)
	}
}

func TestPeriodError_Error(t *testing.T) {
	err := &PeriodError{
		Start: TimeForTest("11:00"),
		End:   TimeForTest("09:00"),
		Err:   ErrPeriodValueEndIsBeforeStart,
	}

	got := err.Error()

	want := "invalid period value `11:00 - 09:00`: end is before start"
	if got != want {
		t.Errorf("PeriodError.Error() = %v, want %v", got, want)
	}
}

func TestPeriodError_Unwrap(t *testing.T) {
	err := &PeriodError{Err: ErrPeriodValueEndIsBeforeStart}

	got := err.Unwrap()

	want := ErrPeriodValueEndIsBeforeStart
	if got != want {
		t.Errorf("PeriodError.Unwrap()= %v, want= %v", got, want)
	}
}
