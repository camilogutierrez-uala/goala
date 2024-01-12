package usrv

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Builder[I any, O any] struct {
	middlewares         []Middleware[I, O]
	server              Server[I, O]
	adapter             AdapterFn[I]
	decorateInterceptor func(any) any
}

func NewBuilder[I any, O any](srv Server[I, O]) *Builder[I, O] {
	return &Builder[I, O]{
		server: srv,
	}
}

func (b *Builder[I, O]) WithMiddlewares(middlewares ...Middleware[I, O]) *Builder[I, O] {
	if len(middlewares) > 0 {
		b.middlewares = append(b.middlewares, middlewares...)
	}
	return b
}

func (b *Builder[I, O]) WithAdapter(fn AdapterFn[I]) *Builder[I, O] {
	b.adapter = fn
	return b
}

func (b *Builder[I, O]) Handler() *Handler[I, O] {
	return NewHandler(
		b.server,
		b.middlewares...,
	)
}

func (b *Builder[I, O]) Serve() {
	if b.adapter == nil {
		b.adapter = AdaptJSON[I]
	}

	interceptor := Interceptor(
		b.Handler(),
		b.adapter,
	)

	if b.decorateInterceptor == nil {
		b.decorateInterceptor = func(a any) any {
			return a
		}
	}

	lambda.Start(
		b.decorateInterceptor(interceptor),
	)
}
