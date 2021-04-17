package services

import (
	"context"
	"fmt"

	"oto-showcase/generated"
)

type ErrorService struct {
}

func (e ErrorService) Err(ctx context.Context, request generated.ErrorRequest) (*generated.ErrorResponse, error) {
	return nil, fmt.Errorf("error example")
}
