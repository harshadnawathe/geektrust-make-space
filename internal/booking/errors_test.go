package booking

import "testing"

func TestInvalidRequestTypeError_Error(t *testing.T) {
	err := &InvalidRequestTypeError{
		Request:      "some string",
		ExpectedType: BookRoomRequest{},
		Err:          ErrInvalidRequestType,
	}

	want := "request type must be `booking.BookRoomRequest`, given `string`: invalid request type"

	if got := err.Error(); got != want {
		t.Errorf("InvalidRequestTypeError.Error() = %v, want %v", got, want)
	}
}

func TestInvalidRequestTypeError_Unwrap(t *testing.T) {
	err := &InvalidRequestTypeError{
		Err:          ErrInvalidRequestType,
	}

	got := err.Unwrap()

	wantErr := ErrInvalidRequestType
	if got != wantErr {
		t.Errorf("InvalidRequestTypeError.Unwrap() error = %v, wantErr %v", err, wantErr)
	}
}

func Test_newInvalidRequestTypeError(t *testing.T) {
	err := newInvalidRequestTypeError("a", "b")

	wantRequest := "a"
	if got := err.Request; got != wantRequest {
    t.Errorf("newInvalidRequestTypeError() = %v, want %v", got, wantRequest)
  }

	wantExpectedType := "b"
  if got := err.ExpectedType; got != wantExpectedType {
		t.Errorf("newInvalidRequestTypeError() = %v, want %v", got, wantExpectedType)
	}

	wantErr := ErrInvalidRequestType
	if got := err.Err; got != wantErr {
    t.Errorf("newInvalidRequestTypeError() = %v, want %v", got, wantErr)
  }
}
