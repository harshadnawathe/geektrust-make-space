package booking_test

import (
	"context"
	"errors"
	"geektrust/internal/booking"
	"geektrust/internal/booking/mocks"
	"geektrust/internal/workplace"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestBookRoomEndpoint(t *testing.T) {
	type args struct {
		request interface{}
	}

	tests := []struct {
		name       string
		args       args
		mockConfig func(*mocks.MockBooker, args)
		want       interface{}
		wantErr    error
	}{
		{
			name: "return response with Reservation",
			args: args{
				request: booking.BookRoomRequest{
					NumOfPeople: 4,
					Period: booking.NewPeriodMust(workplace.NewPeriod(
						booking.NewTimeMust(workplace.NewTime(10, 0)),
						booking.NewTimeMust(workplace.NewTime(10, 30)),
					)),
				},
			},
			mockConfig: func(mock *mocks.MockBooker, args args) {
				req := args.request.(booking.BookRoomRequest)

				mock.EXPECT().Book(req.Period, req.NumOfPeople).
					Times(1).
					Return(workplace.Reservation{Room: "C-Cave"}, nil)
			},
			want:    booking.BookRoomResponse{Reservation: workplace.Reservation{Room: "C-Cave"}, Err: nil},
			wantErr: nil,
		},
		{
			name: "return response with Err",
			args: args{
				request: booking.BookRoomRequest{
					NumOfPeople: 4,
					Period: booking.NewPeriodMust(workplace.NewPeriod(
						booking.NewTimeMust(workplace.NewTime(10, 0)),
						booking.NewTimeMust(workplace.NewTime(10, 30)),
					)),
				},
			},
			mockConfig: func(mock *mocks.MockBooker, args args) {
				req := args.request.(booking.BookRoomRequest)

				mock.EXPECT().Book(req.Period, req.NumOfPeople).
					Times(1).
					Return(workplace.Reservation{}, workplace.ErrRoomNoVacantRoom)
			},
			want:    booking.BookRoomResponse{Reservation: workplace.Reservation{}, Err: workplace.ErrRoomNoVacantRoom},
			wantErr: nil,
		},
		{
			name: "return error when request type is not valid",
			args: args{
				request: struct{ Foo, Bar string }{"foo", "bar"},
			},
			mockConfig: nil,
			want:       nil,
			wantErr:    booking.ErrInvalidRequestType,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			booker := mocks.NewMockBooker(ctrl)
			if test.mockConfig != nil {
				test.mockConfig(booker, test.args)
			}

			endpoint := booking.MakeBookRoomEndpoint(booker)

			got, gotErr := endpoint(context.Background(), test.args.request)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("BookRoom Endpoint()= %v, want= %v", got, test.want)
			}

			if !errors.Is(gotErr, test.wantErr) {
				t.Errorf("BookRoom Endpoint() error = %v, wantErr %v", gotErr, test.wantErr)
			}
		})
	}
}

func TestBookRoomEndpointError_Error(t *testing.T) {
	err := &booking.BookRoomEndpointError{
		Request: struct{ Foo, Bar string }{"foo", "bar"},
		Err:     errors.New("foo"),
	}

	got := err.Error()

	want := "cannot handle request `{foo bar}`: foo"
	if got != want {
		t.Errorf("BookRoomEndpointError.Error() = %v, want %v", got, want)
	}
}

func TestBookRoomEndpointError_Unwrap(t *testing.T) {
	err := &booking.BookRoomEndpointError{Err: booking.ErrInvalidRequestType}

	got := err.Unwrap()

	want := booking.ErrInvalidRequestType
	if !errors.Is(got, want) {
		t.Errorf("BookRoomEndpointError.Unwrap() = %v, want %v", got, want)
	}
}
