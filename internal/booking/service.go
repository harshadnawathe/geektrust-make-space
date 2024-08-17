package booking

import (
	"context"
	"errors"
	"fmt"
	"geektrust/internal/workplace"
)

var (
	ErrInvalidRequestType                       = errors.New("invalid request type")
	ErrBookRoomEndpointInvalidRequestType       = fmt.Errorf("request type must be %T: %w", BookRoomRequest{}, ErrInvalidRequestType)
	ErrRoomsAvailableEndpointInvalidRequestType = fmt.Errorf("request type must be %T: %w", RoomsAvailableRequest{}, ErrInvalidRequestType)
)

type EndpointFunc func(context.Context, interface{}) (interface{}, error)

// BookRoom service endpoint

type BookRoomRequest struct {
	Period      workplace.Period
	NumOfPeople workplace.NumOfPeople
}

type BookRoomResponse struct {
	Err         error
	Reservation workplace.Reservation
}

type Booker interface {
	Book(workplace.Period, workplace.NumOfPeople) (workplace.Reservation, error)
}

func MakeBookRoomEndpoint(booker Booker) EndpointFunc {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		bookRoomRequest, err := getBookRoomRequest(request)
		if err != nil {
			return nil, &BookRoomEndpointError{request, err}
		}
		reservation, err := booker.Book(
			bookRoomRequest.Period,
			bookRoomRequest.NumOfPeople,
		)

		return BookRoomResponse{
			Reservation: reservation,
			Err:         err,
		}, nil
	}
}

func getBookRoomRequest(request interface{}) (BookRoomRequest, error) {
	bookRoomRequest, ok := request.(BookRoomRequest)
	if !ok {
		return BookRoomRequest{}, ErrBookRoomEndpointInvalidRequestType
	}
	return bookRoomRequest, nil
}

type BookRoomEndpointError struct {
	Request interface{}
	Err     error
}

func (err *BookRoomEndpointError) Error() string {
	return fmt.Sprintf("cannot handle request `%v`: %s", err.Request, err.Err)
}

func (err *BookRoomEndpointError) Unwrap() error {
	return err.Err
}


// RoomsAvailable service endpoint

type RoomsAvailableRequest struct {
	Period workplace.Period
}

type RoomsAvailableResponse struct {
	Vacancies []workplace.Vacancy
}

type RoomsAvailabler interface {
	RoomsAvailable(workplace.Period) []workplace.Vacancy
}

func MakeRoomsAvailableEndpoint(roomsAvailabler RoomsAvailabler) EndpointFunc {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		roomsAvailableRequest, err := getRoomsAvailableRequest(request)
		if err != nil {
			return nil, &RoomsAvailableEndpointError{request, err}
		}

		vacancies := roomsAvailabler.RoomsAvailable(
			roomsAvailableRequest.Period,
		)

		return RoomsAvailableResponse{
			Vacancies: vacancies,
		}, nil
	}
}

func getRoomsAvailableRequest(request interface{}) (RoomsAvailableRequest, error) {
	roomsAvailableRequest, ok := request.(RoomsAvailableRequest)
	if !ok {
		return RoomsAvailableRequest{}, ErrRoomsAvailableEndpointInvalidRequestType
	}
	return roomsAvailableRequest, nil
}

type RoomsAvailableEndpointError struct {
	Request interface{}
	Err     error
}

func (err *RoomsAvailableEndpointError) Error() string {
	return fmt.Sprintf("cannot handle request `%v`: %s", err.Request, err.Err)
}

func (err *RoomsAvailableEndpointError) Unwrap() error {
	return err.Err
}
