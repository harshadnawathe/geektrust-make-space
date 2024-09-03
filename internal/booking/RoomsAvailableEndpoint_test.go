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

func TestMakeRoomsAvailableEndpoint(t *testing.T) {
	type args struct {
		request    interface{}
		mockConfig func(*mocks.MockRoomsAvailabler, args)
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{
			name: "return vacancies",
			args: args{
				request: booking.RoomsAvailableRequest{
					Period: booking.NewPeriodMust(workplace.NewPeriod(
						booking.NewTimeMust(workplace.NewTime(10, 0)),
						booking.NewTimeMust(workplace.NewTime(10, 30)),
					)),
				},
				mockConfig: func(mock *mocks.MockRoomsAvailabler, args args) {
					mock.EXPECT().
						RoomsAvailable(args.request.(booking.RoomsAvailableRequest).Period).
						Times(1).
						Return([]workplace.Vacancy{{"C-Cave"}})
				},
			},
			want: booking.RoomsAvailableResponse{
				Vacancies: []workplace.Vacancy{{"C-Cave"}},
			},
			wantErr: nil,
		},
		{
			name: "return error",
			args: args{
				request: struct{ Foo, Bar string }{"foo", "bar"},
				mockConfig: func(mock *mocks.MockRoomsAvailabler, args args) {
				},
			},
			want:    nil,
			wantErr: booking.ErrInvalidRequestType,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			availabler := mocks.NewMockRoomsAvailabler(ctrl)
			test.args.mockConfig(availabler, test.args)

			endpoint := booking.MakeRoomsAvailableEndpoint(availabler)

			got, gotErr := endpoint(context.Background(), test.args.request)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("RoomsAvailable Endpoint()= %v, want= %v", got, test.want)
			}

			if !errors.Is(gotErr, test.wantErr) {
				t.Errorf("RoomsAvailable Endpoint() error = %v, wantErr %v", gotErr, test.wantErr)
			}
		})
	}
}

func TestRoomsAvailableEndpointError_Error(t *testing.T) {
	err := &booking.RoomsAvailableEndpointError{
		Request: struct{ Foo, Bar string }{"foo", "bar"},
		Err:     errors.New("foo"),
	}

	got := err.Error()

	want := "cannot handle request `{foo bar}`: foo"
	if got != want {
		t.Errorf("RoomsAvailableEndpointError.Error() = %v, want %v", got, want)
	}
}

func TestRoomsAvailableEndpointError_Unwrap(t *testing.T) {
	err := &booking.RoomsAvailableEndpointError{Err: booking.ErrInvalidRequestType}

	got := err.Unwrap()

	want := booking.ErrInvalidRequestType
	if !errors.Is(got, want) {
		t.Errorf("RoomsAvailableEndpointError.Unwrap() = %v, want %v", got, want)
	}
}
