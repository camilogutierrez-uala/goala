package usrv

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"log"
)

var (
	lambdaOptions []otellambda.Option
	tracer        trace.Tracer
	shutdownFn    []func() error
	ctx           = context.Background()
)

func UseTrace() error {
	tp, err := xrayconfig.NewTracerProvider(ctx)
	if err != nil {
		return err
	}

	otel.SetTracerProvider(tp)

	shutdownFn = append(
		shutdownFn,
		func() error {
			return tp.Shutdown(ctx)
		},
	)

	return err
}

func OTELShutdown() {
	for _, shutdown := range shutdownFn {
		if err := shutdown(); err != nil {
			log.Printf(err.Error())
		}
	}
}

func OTELLambdaOptions() []otellambda.Option {
	if len(lambdaOptions) == 0 {
		log.Fatal("lambda options are empty")
	}

	return lambdaOptions
}

func SetTrancer(tr trace.Tracer) {
	tracer = tr
}

func Tracer() trace.Tracer {
	if tracer == nil {
		tracer = otel.Tracer("default")
	}
	return tracer
}
