package booking

import (
	"context"
	"errors"
	"geektrust/internal/booking/mocks"
	"geektrust/internal/workplace"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func newPeriodMust(p workplace.Period, err error) workplace.Period {
	if err != nil {
		panic(err)
	}
	return p
}

func newTimeMust(p workplace.Time, err error) workplace.Time {
	if err != nil {
		panic(err)
	}
	return p
}

func TestMakeBookRoomEndpoint(t *testing.T) {
	type args struct {
		request    interface{}
		mockConfig func(*mocks.MockBooker, args)
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{
			name: "return response with Reservation",
			args: args{
				request: BookRoomRequest{
					NumOfPeople: 4,
					Period: newPeriodMust(workplace.NewPeriod(
						newTimeMust(workplace.NewTime(10, 0)),
						newTimeMust(workplace.NewTime(10, 30)),
					)),
				},
				mockConfig: func(mock *mocks.MockBooker, args args) {
					req := args.request.(BookRoomRequest)

					mock.EXPECT().Book(req.Period, req.NumOfPeople).
						Times(1).
						Return(workplace.Reservation{Room: "C-Cave"}, nil)
				},
			},
			want:    BookRoomResponse{Reservation: workplace.Reservation{Room: "C-Cave"}, Err: nil},
			wantErr: nil,
		},
		{
			name: "return response with Err",
			args: args{
				request: BookRoomRequest{
					NumOfPeople: 4,
					Period: newPeriodMust(workplace.NewPeriod(
						newTimeMust(workplace.NewTime(10, 0)),
						newTimeMust(workplace.NewTime(10, 30)),
					)),
				},
				mockConfig: func(mock *mocks.MockBooker, args args) {
					req := args.request.(BookRoomRequest)

					mock.EXPECT().Book(req.Period, req.NumOfPeople).
						Times(1).
						Return(workplace.Reservation{}, workplace.ErrRoomNoVacantRoom)
				},
			},
			want:    BookRoomResponse{Reservation: workplace.Reservation{}, Err: workplace.ErrRoomNoVacantRoom},
			wantErr: nil,
		},
		{
			name: "return error when request type is not valid",
			args: args{
				request: struct{ Foo, Bar string }{"foo", "bar"},
				mockConfig: func(_ *mocks.MockBooker, _ args) {
				},
			},
			want:    nil,
			wantErr: ErrBookRoomEndpointInvalidRequestType,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			booker := mocks.NewMockBooker(ctrl)
			test.args.mockConfig(booker, test.args)

			endpoint := MakeBookRoomEndpoint(booker)

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
				request: RoomsAvailableRequest{
					Period: newPeriodMust(workplace.NewPeriod(
						newTimeMust(workplace.NewTime(10, 0)),
						newTimeMust(workplace.NewTime(10, 30)),
					)),
				},
				mockConfig: func(mock *mocks.MockRoomsAvailabler, args args) {
					mock.EXPECT().
						RoomsAvailable(args.request.(RoomsAvailableRequest).Period).
						Times(1).
						Return([]workplace.Vacancy{{"C-Cave"}})
				},
			},
			want: RoomsAvailableResponse{
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
			wantErr: ErrRoomsAvailableEndpointInvalidRequestType,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			availabler := mocks.NewMockRoomsAvailabler(ctrl)
			test.args.mockConfig(availabler, test.args)

			endpoint := MakeRoomsAvailableEndpoint(availabler)

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

func TestBookRoomEndpointError_Error(t *testing.T) {
	err := &BookRoomEndpointError{
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
	err := &BookRoomEndpointError{Err: ErrInvalidRequestType}

	got := err.Unwrap()

	want := ErrInvalidRequestType
	if !errors.Is(got, want) {
		t.Errorf("BookRoomEndpointError.Unwrap() = %v, want %v", got, want)
	}
}

func TestRoomsAvailableEndpointError_Error(t *testing.T) {
	err := &RoomsAvailableEndpointError{
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
	err := &RoomsAvailableEndpointError{Err: ErrInvalidRequestType}

	got := err.Unwrap()

	want := ErrInvalidRequestType
	if !errors.Is(got, want) {
		t.Errorf("RoomsAvailableEndpointError.Unwrap() = %v, want %v", got, want)
	}
}
