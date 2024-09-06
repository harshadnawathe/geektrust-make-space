package cli

import (
	"context"
	"io"
)

type Handler interface {
	Handle(context.Context, io.Writer, string)
}

type HandlerFunc func(context.Context, io.Writer, string)

func (f HandlerFunc) Handle(ctx context.Context, w io.Writer, cmd string) {
	f(ctx, w, cmd)
}
