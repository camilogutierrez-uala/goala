package usrv

import (
	"context"
)

type (
	Service[I any, O any] func(ctx context.Context, in *I) (*O, error)

	Middleware[I any, O any] func(next Service[I, O]) Service[I, O]

	Handler[I any, O any] struct {
		srv     func(ctx context.Context, in *I) (*O, error)
		isBatch bool
	}
)

func NewHandler[I any, O any](
	srv Service[I, O],
	middlewares ...Middleware[I, O],
) *Handler[I, O] {

	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		srv = middlewares[i](srv)
	}
	return &Handler[I, O]{
		srv: srv,
	}
}

func (h *Handler[I, O]) Batch() {
	h.isBatch = true
}

func (h *Handler[I, O]) EventHandlerWithResponse(ctx context.Context, in *Event[I, O]) (any, error) {
	in.Use(h.srv)
	return in.Process(ctx, h.isBatch)
}

func (h *Handler[I, O]) EventHandlerWithOnlyError(ctx context.Context, in *Event[I, O]) error {
	in.Use(h.srv)
	_, err := in.Process(ctx, false)
	return err
}
