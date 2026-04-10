package main

import (
	"context"
	"net/http"

	"github.com/cylixlee/go-framework-playground/internal/impl"
	"github.com/cylixlee/go-framework-playground/svc"
	"github.com/cylixlee/go-framework-playground/transport"
	"github.com/go-kit/kit/endpoint"
	httpkit "github.com/go-kit/kit/transport/http"
)

func main() {
	srv := new(impl.AddServiceImpl)
	sumEndpoint := httpkit.NewServer(makeSumEndpoint(srv), transport.DecodeSumRequest, transport.EncodeSumResponse)
	concatEndpoint := httpkit.NewServer(makeConcatEndpoint(srv), transport.DecodeConcatRequest, transport.EncodeConcatResponse)

	http.Handle("/sum", sumEndpoint)
	http.Handle("/concat", concatEndpoint)
	http.ListenAndServe(":8080", nil)
}

func makeSumEndpoint(addService svc.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(transport.SumRequest)
		v, err := addService.Sum(ctx, req.A, req.B)
		if err != nil {
			return transport.SumResponse{Err: err}, nil
		}
		return transport.SumResponse{V: v}, nil
	}
}

func makeConcatEndpoint(addService svc.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(transport.ConcatRequest)
		v, err := addService.Concat(ctx, req.A, req.B)
		if err != nil {
			return transport.ConcatResponse{Err: err}, nil
		}
		return transport.ConcatResponse{V: v}, nil
	}
}
