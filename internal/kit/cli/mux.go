package cli

import (
	"context"
	"io"
	"regexp"
)

type CommandMux struct {
	Default  Handler
	handlers []patternHandler
}

type patternHandler struct {
	pattern *regexp.Regexp
	handler Handler
}

func (mux *CommandMux) HandlePattern(pattern *regexp.Regexp, handler Handler) {
	mux.handlers = append(mux.handlers, patternHandler{pattern, handler})
}

func (mux *CommandMux) HandlePatternFunc(pattern *regexp.Regexp, f func(context.Context, io.Writer, string)) {
	mux.HandlePattern(pattern, HandlerFunc(f))
}

func (mux *CommandMux) Handle(ctx context.Context, w io.Writer, cmd string) {
	for _, ph := range mux.handlers {
		if ph.pattern.MatchString(cmd) {
			ph.handler.Handle(ctx, w, cmd)
			return
		}
	}

	if mux.Default != nil {
		mux.Default.Handle(ctx, w, cmd)
	}
}
