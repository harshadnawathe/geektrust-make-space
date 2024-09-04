package kit

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestEndpointFunc_InvokesFunc(t *testing.T) {
	someErr := errors.New("some error")

	type userKey struct{}

	type args struct {
		ctx context.Context
		req interface{}
	}
	tests := []struct {
		name    string
		f       EndpointFunc
		args    args
		want    interface{}
		wantErr error
	}{
		{
			name: "invoke func and return response",
			f: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				user := ctx.Value(userKey{}).(string)

				return fmt.Sprintf("%s, %s", strings.ToUpper(req.(string)), user), nil
			},
			args: args{
				ctx: context.WithValue(context.Background(), userKey{}, "Abc"),
				req: "hello",
			},
			want:    "HELLO, Abc",
			wantErr: nil,
		},
		{
			name: "invoke func and return error",
			f: func(ctx context.Context, req interface{}) (res interface{}, err error) {
				return nil, someErr
			},
			args: args{
				req: "hello",
			},
			want:    nil,
			wantErr: someErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.f.Invoke(tt.args.ctx, tt.args.req)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("EndpointFunc.Invoke() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndpointFunc.Invoke() = %v, want %v", got, tt.want)
			}
		})
	}
}
