package service

import (
	"context"
)

type (
	Service struct {
	}

	Request struct {
		PaymentMethod string
		Amount        int64
	}

	Response struct {
		Status   string
		Terminal string
	}
)

func New() *Service {
	return &Service{}
}

func (a *Service) Service(ctx context.Context, in *Request) (*Response, error) {
	return &Response{}, nil
}
