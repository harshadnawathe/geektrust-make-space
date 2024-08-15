package workplace

import (
	"errors"
	"reflect"
	"testing"
)

func Test_newRoom(t *testing.T) {
	type args struct {
		name     string
		capacity NumOfPeople
	}
	tests := []struct {
		name    string
		args    args
		want    *room
		wantErr error
	}{
		{
			name: "create room with no error",
			args: args{
				name:     "C-Cave",
				capacity: 3,
			},
			want:    &room{"C-Cave", nil, 3},
			wantErr: nil,
		},
		{
			name: "errow when name is blank",
			args: args{
				name:     " ",
				capacity: 3,
			},
			want:    nil,
			wantErr: ErrRoomNameIsBlank,
		},
		{
			name: "error when capacity is zero",
			args: args{
				name:     "C-Cave",
				capacity: 0,
			},
			want:    nil,
			wantErr: ErrRoomCapacityIsZero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newRoom(tt.args.name, tt.args.capacity)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRoom() = %v, want %v", got, tt.want)
			}

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("newRoom() error = %v, wantErr %v", err, tt.wantErr)
				}

				var roomInitErr *RoomInitError
				if !errors.As(err, &roomInitErr) {
					t.Errorf("cannot use error %v as %T", err, roomInitErr)
				}

				if roomInitErr.Name != tt.args.name {
					t.Errorf("RoomInitError.Name = %v, want %v", roomInitErr.Name, tt.args.name)
				}

				if roomInitErr.Capacity != tt.args.capacity {
					t.Errorf("RoomInitError.Capacity = %v, want %v", roomInitErr.Capacity, tt.args.capacity)
				}

				return
			}
		})
	}
}

func TestRoomInitError_Error(t *testing.T) {
	err := &RoomInitError{
		Name:     "C-Cave",
		Capacity: 45,
		Err:      errors.New("some error"),
	}

	got := err.Error()

	want := "cannot create a room with name `C-Cave` and capacity `45`: some error"
	if got != want {
		t.Errorf("RoomInitError.Error() = %v, want %v", got, want)
	}
}

func TestRoomInitError_Unwrap(t *testing.T) {
	err := &RoomInitError{
		Err: ErrRoomCapacityIsZero,
	}

	got := err.Unwrap()

	want := ErrRoomCapacityIsZero
	if got != want {
		t.Errorf("RoomInitError.Unwrap() = %v, want %v", got, want)
	}
}

func TestRoomReserveError_Error(t *testing.T) {
	err := &RoomReserveError{
		Period:      PeriodForTest("09:00", "10:00"),
		NumOfPeople: 5,
		Err:         errors.New("some error"),
	}

	got := err.Error()

	want := "cannot reserve room for `5` people in period `09:00 - 10:00`: some error"
	if got != want {
		t.Errorf("RoomReserveError.Error() = %v, want %v", got, want)
	}
}

func TestRoomReserveError_Unwrap(t *testing.T) {
	err := &RoomReserveError{
		Err: ErrRoomNoVacantRoom,
	}

	got := err.Unwrap()

	want := ErrRoomNoVacantRoom
	if got != want {
		t.Errorf("RoomReserveError.Unwrap() = %v, want %v", got, want)
	}
}

func TestRoomCapacityValidationError_Error(t *testing.T) {
	err := &RoomCapacityValidationError{
		Name:        "C-Cave",
		Capacity:    2,
		NumOfPeople: 4,
		Err:         errors.New("some error"),
	}

	got := err.Error()

	want := "cannot fit `4` people in room `C-Cave` with capacity `2`: some error"
	if got != want {
		t.Errorf("RoomCapacityValidationError.Error() = %v, want %v", got, want)
	}
}

func TestRoomCapacityValidationError_Unwrap(t *testing.T) {
	err := &RoomCapacityValidationError{Err: ErrRoomCapacityIsTooSmall}

	got := err.Unwrap()

	want := ErrRoomCapacityIsTooSmall

	if got != want {
		t.Errorf("RoomCapacityValidationError.Unwrap() = %v, want %v", got, want)
	}
}
