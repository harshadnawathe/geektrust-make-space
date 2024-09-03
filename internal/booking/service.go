package booking

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRequestType                       = errors.New("invalid request type")
	ErrBookRoomEndpointInvalidRequestType       = fmt.Errorf("request type must be %T: %w", BookRoomRequest{}, ErrInvalidRequestType)
	ErrRoomsAvailableEndpointInvalidRequestType = fmt.Errorf("request type must be %T: %w", RoomsAvailableRequest{}, ErrInvalidRequestType)
)

type InvalidRequestTypeError struct {
	Request, ExpectedType interface{}
	Err                   error
}

func (err *InvalidRequestTypeError) Error() string {
	return fmt.Sprintf(
		"request type must be `%T`, given `%T`: %s",
		err.ExpectedType,
		err.Request,
		ErrInvalidRequestType,
	)
}

func (err *InvalidRequestTypeError) Unwrap() error {
	return err.Err
}

func newInvalidRequestTypeError(req, exp interface{}) error {
	return &InvalidRequestTypeError{req, exp, ErrInvalidRequestType}
}
