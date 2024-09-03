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

