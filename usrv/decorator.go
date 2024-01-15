package usrv

import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"

func OTELDecorator(opts ...otellambda.Option) func(any) any {
	return func(a any) any {
		return otellambda.InstrumentHandler(a, opts...)
	}
}
