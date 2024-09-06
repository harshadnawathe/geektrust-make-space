package main

import (
	"bytes"
	"context"
	"geektrust/internal/workplace"
	"testing"
)

func TestMakeVacancyCommandHandler(t *testing.T) {
	type args struct {
		ctx context.Context
		cmd string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "write INCORRECT_INPUT when command doesn't match",
			args: args{
				cmd: "BOOK 14:00 15:30 3",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when start hh is incorrect",
			args: args{
				cmd: "VACANCY 30:00 12:00",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when start mm is incorrect",
			args: args{
				cmd: "VACANCY 10:70 12:00",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end hh is incorrect",
			args: args{
				cmd: "VACANCY 10:00 42:00",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end mm is incorrect",
			args: args{
				cmd: "VACANCY 10:00 12:88",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end is before start",
			args: args{
				cmd: "VACANCY 12:00 10:00",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write NO_VACANT_ROOMS when no rooms are available to book",
			args: args{
				cmd: "VACANCY 10:30 11:00",
			},
			wantW: "NO_VACANT_ROOM\n",
		},
		{
			name: "write room names when rooms are available to book",
			args: args{
				cmd: "VACANCY 12:00 12:30",
			},
			wantW: "C-CAVE D-TOWER\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := workplace.New()

			_ = wp.AddRoom("C-CAVE", 3)
			_ = wp.AddRoom("D-TOWER", 7)

			p1, _ := parsePeriod("10:00", "11:00")
			_, _ = wp.Book(p1, 2)

			p2, _ := parsePeriod("10:30", "11:30")
			_, _ = wp.Book(p2, 2)

			cmdHandler := MakeVacancyCommandHandler(wp)

			w := &bytes.Buffer{}

			cmdHandler(tt.args.ctx, w, tt.args.cmd)

			gotW := w.String()
			if gotW != tt.wantW {
				t.Errorf("MakeVacancyCommandHandler() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
