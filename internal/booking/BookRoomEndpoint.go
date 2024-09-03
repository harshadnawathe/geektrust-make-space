package booking

import (
	"context"
	"fmt"
	"geektrust/internal/workplace"
)

// BookRoom service endpoint
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
