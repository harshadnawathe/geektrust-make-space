package main

import (
	"context"
	"fmt"
	"geektrust/internal/workplace"
	"io"
	"regexp"
	"strings"
)


var vacancyCommandPattern = regexp.MustCompile(`^VACANCY\s+(\d\d:\d\d)\s+(\d\d:\d\d)$`)

func MakeVacancyCommandHandler(wp *workplace.Workplace) func(context.Context, io.Writer, string) {
	return func(ctx context.Context, w io.Writer, cmd string) {
		tokens := vacancyCommandPattern.FindStringSubmatch(cmd)
		if len(tokens) == 0 {
			_, _ = fmt.Fprintln(w, incorrectInput)
			return
		}

		period, err := parsePeriod(tokens[1], tokens[2])
		if err != nil {
			_, _ = fmt.Fprintln(w, incorrectInput)
			return
		}

		vacancies := wp.RoomsAvailable(period)
		if len(vacancies) == 0 {
			_, _ = fmt.Fprintln(w, noVacantRoom)
			return
		}

		rooms := make([]string, 0, len(vacancies))
		for _, vacancy := range vacancies {
			rooms = append(rooms, vacancy.Room)
		}

		_, _ = fmt.Fprintln(w, strings.Join(rooms, " "))
	}
}
