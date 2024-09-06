package main

import (
	"bytes"
	"context"
	"geektrust/internal/workplace"
	"testing"
)

func TestMakeBookCommandHandler(t *testing.T) {
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
			args: args {
				cmd: "VACANCY 12:00 12:30",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when start hh is incorrect",
			args: args{
				cmd: "BOOK 30:00 12:00 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when start mm is incorrect",
			args: args{
				cmd: "BOOK 10:70 12:00 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end hh is incorrect",
			args: args{
				cmd: "BOOK 10:00 42:00 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end mm is incorrect",
			args: args{
				cmd: "BOOK 10:00 12:88 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end is before start",
			args: args{
				cmd: "BOOK 12:00 10:00 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when start minutes is not in 15 minutes increments",
			args: args{
				cmd: "BOOK 12:24 10:00 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write INCORRECT_INPUT when end minutes is not in 15 minutes increments",
			args: args{
				cmd: "BOOK 12:00 10:04 2",
			},
			wantW: "INCORRECT_INPUT\n",
		},
		{
			name: "write NO_VACANT_ROOM when period overlaps with buffer time",
			args: args{
				cmd: "BOOK 08:30 09:30 2",
			},
			wantW: "NO_VACANT_ROOM\n",
		},
		{
			name: "write NO_VACANT_ROOM when no room is available",
			args: args{
				cmd: "BOOK 09:30 10:30 20",
			},
			wantW: "NO_VACANT_ROOM\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := workplace.New()

			p, _ := parsePeriod("09:00", "09:15")
			wp.AddBufferTime(p)

			_ = wp.AddRoom("C-CAVE", 3)
			_ = wp.AddRoom("D-TOWER", 7)

			cmdHandler := MakeBookCommandHandler(wp)

			w := &bytes.Buffer{}

			cmdHandler(tt.args.ctx, w, tt.args.cmd)

			gotW := w.String()
			if gotW != tt.wantW {
				t.Errorf("BookCommandHandler() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
