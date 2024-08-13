package workplace

import (
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
		wantErr bool
	}{
		{
			name: "create time with no error",
			args: args{
				hh: 11,
				mm: 59,
			},
			want:    Time{11, 59},
			wantErr: false,
		},
		{
			name: "error when hour value is greater than 23",
			args: args{
				hh: 24,
				mm: 59,
			},
			want:    Time{},
			wantErr: true,
		},
		{
			name: "error when minute value is greater than 59",
			args: args{
				hh: 12,
				mm: 60,
			},
			want:    Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTime(tt.args.hh, tt.args.mm)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
