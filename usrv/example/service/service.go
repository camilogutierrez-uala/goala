package service

import (
	"context"
	"github.com/camilogutierrez-uala/goala/usrv"
)

type (
	service struct {
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

func New() usrv.Server[Request, Response] {
	return &service{}
}

func (a *service) Service(ctx context.Context, in *Request) (*Response, error) {
	return &Response{}, nil
}
