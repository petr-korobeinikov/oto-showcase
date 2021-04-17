package services

import (
	"context"

	"oto-showcase/generated"
)

type FoobarService struct {
}

func (f FoobarService) Bar(ctx context.Context, request generated.BarRequest) (*generated.BarResponse, error) {
	resp := &generated.BarResponse{
		Bar: request.Bar,
	}
	return resp, nil
}

func (f FoobarService) Foo(ctx context.Context, request generated.FooRequest) (*generated.FooResponse, error) {
	resp := &generated.FooResponse{
		IntVal: 1 * request.Multiplier,
		StrVal: "str",
		ArrVal: []string{"foo", "bar", "baz"},
	}
	return resp, nil
}
