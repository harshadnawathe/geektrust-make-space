package main

import (
	"context"
	"fmt"
	"geektrust/internal/kit/cli"
	"geektrust/internal/workplace"
	"io"
)

var defaultHandler = cli.HandlerFunc(func(_ context.Context, w io.Writer, _ string) {
	_, _ = fmt.Fprintln(w, incorrectInput)
})

func MakeWorkplaceCommandHandler(wp *workplace.Workplace) *cli.CommandMux {
	mux := &cli.CommandMux{
		Default: defaultHandler,
	}

	mux.HandlePatternFunc(bookCommandPattern, MakeBookCommandHandler(wp))
	mux.HandlePatternFunc(vacancyCommandPattern, MakeVacancyCommandHandler(wp))

	return mux
}
