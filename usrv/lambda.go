package usrv

func Serve[I any, O any](srv Server[I, O], middleware ...Middleware[I, O]) {
	build := NewBuilder[I, O](
		srv,
	).WithMiddlewares(
		Logger[I, O],
	).WithMiddlewares(
		middleware...,
	)

	build.Serve()
}

func ServeHTTP[I any, O any](srv Server[I, O], middleware ...Middleware[I, O]) {
	build := NewBuilder[I, O](
		srv,
	).WithMiddlewares(
		Logger[I, O],
	).WithMiddlewares(
		middleware...,
	)

	build.ServeHTTP()
}
