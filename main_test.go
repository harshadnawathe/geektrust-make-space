package main

import (
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	os.Stdout = stdout
	_ = w.Close()

	output, _ := io.ReadAll(r)

	return string(output)
}

func Test_main(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "process input 1",
			input: "sample_input/input1.txt",
			want: `C-Cave D-Tower G-Mansion
C-Cave
NO_VACANT_ROOM
G-Mansion
D-Tower
C-Cave
D-Tower
INCORRECT_INPUT
C-Cave
G-Mansion
G-Mansion
NO_VACANT_ROOM
`,
		},
		{
			name:  "process input 2",
			input: "sample_input/input2.txt",
			want: `C-Cave
C-Cave
NO_VACANT_ROOM
D-Tower
G-Mansion
G-Mansion
G-Mansion
NO_VACANT_ROOM
D-Tower
NO_VACANT_ROOM
INCORRECT_INPUT
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{"geektrust", tt.input}

			got := captureStdout(main)

			if got != tt.want {
				t.Errorf("main()= %s want= %s", got, tt.want)
			}
		})
	}
}
