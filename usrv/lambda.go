package usrv

import (
	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

func baseLambda[I any, O any](srv Server[I, O], middleware ...Middleware[I, O]) *Builder[I, O] {
	builder := NewBuilder[I, O](
		srv,
	).WithMiddlewares(
		Logger[I, O],
	).WithMiddlewares(
		middleware...,
	)
	return builder
}

func LambdaServe[I any, O any](srv Server[I, O], middleware ...Middleware[I, O]) {
	lambda.Start(baseLambda(srv, middleware...).Build())
}

func LambdaOTELServe[I any, O any](srv Server[I, O], opts []otellambda.Option, middleware ...Middleware[I, O]) {
	builder := baseLambda(srv, middleware...).
		WithInterceptorDecoration(
			OTELDecorator(opts...),
		)
	lambda.Start(builder.Build())
}
