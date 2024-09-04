package kit

import "context"

type Endpoint interface {
	Invoke(context.Context, interface{}) (interface{}, error)
}

type EndpointFunc func(ctx context.Context, req interface{}) (res interface{}, err error)

func (f EndpointFunc) Invoke(ctx context.Context, req interface{}) (interface{}, error) {
  return f(ctx, req)
}
