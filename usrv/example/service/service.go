package service

import (
	"context"
	"errors"
	"github.com/Bancar/goala/usrv"
	"strings"
	"time"
)

type (
	Service struct {
	}

	Request struct {
		Amount        float64        `dynamodbav:"Amount"`
		Message       string         `dynamodbav:"Message"`
		CardData      []byte         `dynamodbav:"CardData"`
		Available     bool           `dynamodbav:"Available"`
		Null          any            `dynamodbav:"Null"`
		ChangeDetails []any          `dynamodbav:"ChangeDetails"`
		Detail        map[string]any `dynamodbav:"Detal"`
		Terminals     []float64      `dynamodbav:"Terminals"`
		Products      []string       `dynamodbav:"Products"`
		Token         [][]byte       `dynamodbav:"Token"`
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
	if strings.Contains(in.Message, "Hello, fail!") {
		return nil, errors.New("internal service error")
	}

	func() {
		_, span := usrv.Tracer().Start(ctx, "DB")
		defer span.End()

		time.Sleep(1 * time.Millisecond)
	}()
	return &Response{
		Status:   "OK",
		Terminal: "terminal-123",
	}, nil
}
