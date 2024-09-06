package main

import (
	"context"
	"geektrust/internal/workplace"
	"testing"
)

func TestWorkpalceCommandHandler(t *testing.T) {
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
			name: "write rooms names",
			args: args{
				cmd: "VACANCY 09:30 10:30",
			},
			wantW: "C-CAVE D-TOWER\n",
		},
		{
			name: "write rooms name after booking",
			args: args{
				cmd: "BOOK 09:30 10:30 2",
			},
			wantW: "C-CAVE\n",
		},
		{
			name: "write INCORRECT_INPUT",
			args: args{
				cmd: "SOME COMMAND",
			},
			wantW: "INCORRECT_INPUT\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := workplace.New()

			_ = wp.AddRoom("C-CAVE", 3)
			_ = wp.AddRoom("D-TOWER", 7)
		})
	}
}
