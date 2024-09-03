package booking

import (
	"context"
	"geektrust/internal/workplace"
	"reflect"
	"testing"
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

func TestServiceEndpointsIntegration(t *testing.T) {
	w := workplace.New()

	_ = w.AddRoom("C-Cave", 3)
	_ = w.AddRoom("D-Tower", 7)

	w.AddBufferTime(newPeriodMust(workplace.NewPeriod(
		newTimeMust(workplace.NewTime(9, 0)),
		newTimeMust(workplace.NewTime(9, 15)),
	)))

	roomsAvailableEndpoint := MakeRoomsAvailableEndpoint(w)
	bookRoomEndpoint := MakeBookRoomEndpoint(w)

	var err error

	// Test roomsAvailableEndpoint before booking
	var roomsAvailableRequest, roomsAvailableResponse interface{}
	wantRoomsAvailableResponse := RoomsAvailableResponse{
		Vacancies: []workplace.Vacancy{{"C-Cave"}, {"D-Tower"}},
	}

	roomsAvailableRequest = RoomsAvailableRequest{
		Period: newPeriodMust(workplace.NewPeriod(
			newTimeMust(workplace.NewTime(10, 0)),
			newTimeMust(workplace.NewTime(10, 30)),
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

	bookRoomRequest = BookRoomRequest{
		NumOfPeople: 2,
		Period: newPeriodMust(workplace.NewPeriod(
			newTimeMust(workplace.NewTime(10, 0)),
			newTimeMust(workplace.NewTime(11, 0)),
		)),
	}

	wantBookRoomResponse := BookRoomResponse{Reservation: workplace.Reservation{Room: "C-Cave"}, Err: nil}

	bookRoomResponse, err = bookRoomEndpoint(context.Background(), bookRoomRequest)
	if err != nil {
		t.Errorf("BookRoom Endpoint() error = %v", err)
	}

	if !reflect.DeepEqual(bookRoomResponse, wantBookRoomResponse) {
		t.Errorf("BookRoom Endpoint()= %v, want= %v", bookRoomResponse, wantBookRoomResponse)
	}

	// Test roomsAvailableEndpoint after booking
	wantRoomsAvailableResponse = RoomsAvailableResponse{
		Vacancies: []workplace.Vacancy{{"D-Tower"}},
	}

	roomsAvailableRequest = RoomsAvailableRequest{
		Period: newPeriodMust(workplace.NewPeriod(
			newTimeMust(workplace.NewTime(10, 0)),
			newTimeMust(workplace.NewTime(10, 30)),
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
