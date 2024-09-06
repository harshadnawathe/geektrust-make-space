package main

import (
	"context"
	"errors"
	"fmt"
	"geektrust/internal/workplace"
	"io"
	"regexp"
	"strconv"
)

var bookCommandPattern = regexp.MustCompile(`^BOOK\s+(\d\d:\d\d)\s+(\d\d:\d\d)\s+(\d+)$`)

func MakeBookCommandHandler(wp *workplace.Workplace) func(context.Context, io.Writer, string) {
	return func(ctx context.Context, w io.Writer, cmd string) {
		tokens := bookCommandPattern.FindStringSubmatch(cmd)
		if len(tokens) == 0 {
			_, _ = fmt.Fprintln(w, incorrectInput)
			return
		}

		period, err := parsePeriod(tokens[1], tokens[2])
		if err != nil {
			_, _ = fmt.Fprintln(w, incorrectInput)
			return
		}

		personCapacity, err := strconv.Atoi(tokens[3])
		if err != nil {
			_, _ = fmt.Fprintln(w, incorrectInput)
			return
		}

		reservation, err := wp.Book(period, workplace.NumOfPeople(personCapacity))
		
		if errors.Is(err, workplace.ErrRoomNoVacantRoom) || errors.Is(err, workplace.ErrPeriodOverlapsWithBufferTime) {
			_, _ = fmt.Fprintln(w, noVacantRoom)
			return
		}

		if err != nil {
			_, _ = fmt.Fprintln(w, incorrectInput)
			return
		}

		_, _ = fmt.Fprintln(w, reservation.Room)
	}
}
