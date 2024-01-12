package usrv

import (
	"context"
	"encoding/json"
)

type (
	Server[I any, O any] interface {
		Service(ctx context.Context, in *I) (*O, error)
	}

	Middleware[I any, O any] func(next HandlerFn[I, O]) HandlerFn[I, O]
	HandlerFn[I any, O any]  func(ctx context.Context, in *I) (*O, error)
	AdapterFn[I any]         func(raw json.RawMessage) (*I, error)

	Handler[I any, O any] struct {
		middlewares []Middleware[I, O]
		srv         Server[I, O]
	}
)

func NewHandler[I any, O any](
	srv Server[I, O],
	middlewares ...Middleware[I, O],
) *Handler[I, O] {
	return &Handler[I, O]{
		middlewares: middlewares,
		srv:         srv,
	}
}

func (h *Handler[I, O]) EventHandler(ctx context.Context, in *I) (*O, error) {
	next := h.srv.Service
	// Apply middlewares in reverse order
	for i := len(h.middlewares) - 1; i >= 0; i-- {
		next = h.middlewares[i](next)
	}

	return next(ctx, in)
}

func Interceptor[I any, O any](h *Handler[I, O], adapt AdapterFn[I]) func(ctx context.Context, raw json.RawMessage) (*O, error) {
	return func(ctx context.Context, raw json.RawMessage) (*O, error) {
		in, err := adapt(raw)
		if err != nil {
			return nil, err
		}

		return h.EventHandler(ctx, in)
	}
}
