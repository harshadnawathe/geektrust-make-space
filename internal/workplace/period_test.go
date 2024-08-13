package workplace

import "testing"

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
