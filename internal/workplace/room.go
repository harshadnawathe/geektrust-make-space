package workplace

import (
	"errors"
	"fmt"
)

type room struct {
	name     string
	bookedAt []Period
	capacity int
}

func newRoom(name string, capacity int) *room {
	return &room{
		name:     name,
		capacity: capacity,
	}
}

func canFit(r *room, numOfPeople int) bool {
	return numOfPeople <= r.capacity
}

func isBooked(r *room, p Period) bool {
	return isAnyOverlapping(r.bookedAt, p)
}

func book(r *room, p Period, numOfPeople int) error {
	if !canFit(r, numOfPeople) {
		return fmt.Errorf("cannot book: cannot fit %v people in the room with capacity %v", numOfPeople, r.capacity)
	}

	if isBooked(r, p) {
		return errors.New("cannot book: room is booked")
	}

	r.bookedAt = append(r.bookedAt, p)
	return nil
}