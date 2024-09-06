package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"
	"testing"
)

func TestCommandMux_Handle(t *testing.T) {
	var defaultHandler HandlerFunc = func(ctx context.Context, w io.Writer, cmd string) {
		_, _ = fmt.Fprintf(w, "DEFAULT= %s", cmd)
	}

	type patternHandlerFunc struct {
		pattern *regexp.Regexp
		f       func(context.Context, io.Writer, string)
	}

	fooHandler := patternHandlerFunc{
		pattern: regexp.MustCompile(`^FOO\s+`),
		f: func(ctx context.Context, w io.Writer, cmd string) {
			_, _ = fmt.Fprintf(w, "FOO= %s", cmd)
		},
	}

	barHandler := patternHandlerFunc{
		pattern: regexp.MustCompile(`^BAR\s+`),
		f: func(ctx context.Context, w io.Writer, cmd string) {
			_, _ = fmt.Fprintf(w, "BAR= %s", cmd)
		},
	}

	type fields struct {
		Default  Handler
		handlers []patternHandlerFunc
	}

	type args struct {
		ctx context.Context
		cmd string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantW  string
	}{
		{
			name: "use default handling when no handler provided",
			fields: fields{
				Default:  defaultHandler,
				handlers: []patternHandlerFunc{},
			},
			args: args{
				cmd: "SOME COMMAND",
			},
			wantW: "DEFAULT= SOME COMMAND",
		},
		{
			name: "does not use default handling when default handler not provided",
			fields: fields{
				Default:  nil,
				handlers: []patternHandlerFunc{},
			},
			args: args{
				cmd: "SOME COMMAND",
			},
			wantW: "",
		},
		{
			name: "use default handling when handler pattern does not match",
			fields: fields{
				Default:  defaultHandler,
				handlers: []patternHandlerFunc{fooHandler, barHandler},
			},
			args: args{
				cmd: "SOME COMMAND",
			},
			wantW: "DEFAULT= SOME COMMAND",
		},
		{
			name: "use command specific handling when handler pattern matches",
			fields: fields{
				Default:  defaultHandler,
				handlers: []patternHandlerFunc{fooHandler, barHandler},
			},
			args: args{
				cmd: "BAR COMMAND",
			},
			wantW: "BAR= BAR COMMAND",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := &CommandMux{
				Default: tt.fields.Default,
			}

			for _, ph := range tt.fields.handlers {
				mux.HandlePatternFunc(ph.pattern, ph.f)
			}

			w := &bytes.Buffer{}

			mux.Handle(tt.args.ctx, w, tt.args.cmd)

			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("CommandMux.Handle() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
