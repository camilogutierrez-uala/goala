package usrv

type Builder[I any, O any] struct {
	middlewares      []Middleware[I, O]
	service          Service[I, O]
	isErrorOnly      bool
	handlerDecorator func(any) any
}

func NewBuilder[I any, O any](srv Service[I, O]) *Builder[I, O] {
	return &Builder[I, O]{
		service: srv,
	}
}

func (b *Builder[I, O]) WithMiddlewares(middlewares ...Middleware[I, O]) *Builder[I, O] {
	if len(middlewares) > 0 {
		b.middlewares = append(b.middlewares, middlewares...)
	}
	return b
}

func (b *Builder[I, O]) WithOnlyError() *Builder[I, O] {
	b.isErrorOnly = true
	return b
}

func (b *Builder[I, O]) WithHandlerDecorator(fn func(any) any) *Builder[I, O] {
	b.handlerDecorator = fn
	return b
}

func (b *Builder[I, O]) Handler() *Handler[I, O] {
	return NewHandler(
		b.service,
		b.middlewares...,
	)
}

func (b *Builder[I, O]) Build() any {
	var handler any
	if b.isErrorOnly {
		handler = b.Handler().EventHandlerWithOnlyError
	} else {
		handler = b.Handler().EventHandlerWithResponse
	}

	if b.handlerDecorator == nil {
		return handler
	}

	return b.handlerDecorator(handler)
}
