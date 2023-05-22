package endpoint

import (
	"bluebell/gokit/data"
	"bluebell/gokit/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeSumEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(data.SumRequest)
		v, err := svc.Sum(ctx, req.A, req.B)
		if err != nil {
			return data.SumResponse{V: v, Err: err.Error()}, nil
		}
		return data.SumResponse{V: v}, nil
	}
}

func MakeConcatEndpoint(svc service.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(data.ConcatRequest)
		v, err := svc.Concat(ctx, req.A, req.B)
		if err != nil {
			return data.ConcatResponse{V: v, Err: err.Error()}, nil
		}
		return data.ConcatResponse{V: v}, nil
	}
}
