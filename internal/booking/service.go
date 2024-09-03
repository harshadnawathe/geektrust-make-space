package booking

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrInvalidRequestType                       = errors.New("invalid request type")
	ErrBookRoomEndpointInvalidRequestType       = fmt.Errorf("request type must be %T: %w", BookRoomRequest{}, ErrInvalidRequestType)
	ErrRoomsAvailableEndpointInvalidRequestType = fmt.Errorf("request type must be %T: %w", RoomsAvailableRequest{}, ErrInvalidRequestType)
)

type EndpointFunc func(context.Context, interface{}) (interface{}, error)
