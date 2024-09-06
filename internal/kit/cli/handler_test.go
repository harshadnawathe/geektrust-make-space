package cli

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"
)

func TestHandlerFunc_InvokesFunc(t *testing.T) {
	type userKey struct{}

	var h HandlerFunc = func(ctx context.Context, w io.Writer, cmd string) {
		user := ctx.Value(userKey{}).(string)

		_, _ = fmt.Fprintf(w, "%s, %s!", cmd, user)
	}

	ctx := context.WithValue(context.Background(), userKey{}, "Bob")
	buf := bytes.Buffer{}

	h.Handle(ctx, &buf, "Hello")

	want := "Hello, Bob!"
	if got := buf.String(); got != want {
		t.Errorf("HandlerFunc.Handle() = %q, want %q", got, want)
	}
}
