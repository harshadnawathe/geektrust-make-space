package booking_test

import (
	"context"
	"geektrust/internal/booking"
	"geektrust/internal/workplace"
	"reflect"
	"testing"
)

func TestBookingService(t *testing.T) {
	w := workplace.New()

	_ = w.AddRoom("C-Cave", 3)
	_ = w.AddRoom("D-Tower", 7)

	w.AddBufferTime(booking.NewPeriodMust(workplace.NewPeriod(
		booking.NewTimeMust(workplace.NewTime(9, 0)),
		booking.NewTimeMust(workplace.NewTime(9, 15)),
	)))

	roomsAvailableEndpoint := booking.MakeRoomsAvailableEndpoint(w)
	bookRoomEndpoint := booking.MakeBookRoomEndpoint(w)

	var err error

	// Test roomsAvailableEndpoint before booking
	var roomsAvailableRequest, roomsAvailableResponse interface{}
	wantRoomsAvailableResponse := booking.RoomsAvailableResponse{
		Vacancies: []workplace.Vacancy{{"C-Cave"}, {"D-Tower"}},
	}

	roomsAvailableRequest = booking.RoomsAvailableRequest{
		Period: booking.NewPeriodMust(workplace.NewPeriod(
			booking.NewTimeMust(workplace.NewTime(10, 0)),
			booking.NewTimeMust(workplace.NewTime(10, 30)),
		)),
	}

	roomsAvailableResponse, err = roomsAvailableEndpoint(context.Background(), roomsAvailableRequest)
	if err != nil {
		t.Errorf("RoomsAvailable Endpoint() error = %v", err)
	}

	if !reflect.DeepEqual(roomsAvailableResponse, wantRoomsAvailableResponse) {
		t.Errorf("RoomsAvailable Endpoint()= %v, want= %v", roomsAvailableResponse, wantRoomsAvailableResponse)
	}

	// Test bookRoomEndpoint

	var bookRoomRequest, bookRoomResponse interface{}

	bookRoomRequest = booking.BookRoomRequest{
		NumOfPeople: 2,
		Period: booking.NewPeriodMust(workplace.NewPeriod(
			booking.NewTimeMust(workplace.NewTime(10, 0)),
			booking.NewTimeMust(workplace.NewTime(11, 0)),
		)),
	}

	wantBookRoomResponse := booking.BookRoomResponse{Reservation: workplace.Reservation{Room: "C-Cave"}, Err: nil}

	bookRoomResponse, err = bookRoomEndpoint(context.Background(), bookRoomRequest)
	if err != nil {
		t.Errorf("BookRoom Endpoint() error = %v", err)
	}

	if !reflect.DeepEqual(bookRoomResponse, wantBookRoomResponse) {
		t.Errorf("BookRoom Endpoint()= %v, want= %v", bookRoomResponse, wantBookRoomResponse)
	}

	// Test roomsAvailableEndpoint after booking
	wantRoomsAvailableResponse = booking.RoomsAvailableResponse{
		Vacancies: []workplace.Vacancy{{"D-Tower"}},
	}

	roomsAvailableRequest = booking.RoomsAvailableRequest{
		Period: booking.NewPeriodMust(workplace.NewPeriod(
			booking.NewTimeMust(workplace.NewTime(10, 0)),
			booking.NewTimeMust(workplace.NewTime(10, 30)),
		)),
	}

	roomsAvailableResponse, err = roomsAvailableEndpoint(context.Background(), roomsAvailableRequest)
	if err != nil {
		t.Errorf("RoomsAvailable Endpoint() error = %v", err)
	}

	if !reflect.DeepEqual(roomsAvailableResponse, wantRoomsAvailableResponse) {
		t.Errorf("RoomsAvailable Endpoint()= %v, want= %v", roomsAvailableResponse, wantRoomsAvailableResponse)
	}
}
