package usrv

type Builder[I any, O any] struct {
	middlewares          []Middleware[I, O]
	server               Server[I, O]
	adapter              AdapterFn[I]
	interceptorDecorator func(any) any
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

func (b *Builder[I, O]) WithInterceptorDecoration(fn func(any) any) *Builder[I, O] {
	b.interceptorDecorator = fn
	return b
}

func (b *Builder[I, O]) Handler() *Handler[I, O] {
	return NewHandler(
		b.server,
		b.middlewares...,
	)
}

func (b *Builder[I, O]) Build() any {
	if b.adapter == nil {
		b.adapter = AdaptJSON[I]
	}

	interceptor := Interceptor(
		b.Handler(),
		b.adapter,
	)

	if b.interceptorDecorator == nil {
		return interceptor
	}

	return b.interceptorDecorator(interceptor)
}
