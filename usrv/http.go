package usrv

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"io"
	"log"
	"net/http"
	"os"
)

func TraceServiceInvoque[I any, O any](next Service[I, O]) Service[I, O] {
	return func(ctx context.Context, in *I) (*O, error) {
		trx, span := Tracer().Start(ctx, "service.invoke")
		defer span.End()

		res, err := next(trx, in)
		if err != nil {
			span.SetStatus(codes.Error, "service has an error")
			span.RecordError(err)
			return nil, err
		}
		span.SetStatus(codes.Ok, "service execution success")
		return res, nil
	}
}

func LocalHTTP[I any, O any](srv Service[I, O], isBatch bool, middleware ...Middleware[I, O]) {
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	os.Setenv("OTEL_TRACES_EXPORTER", "otlp")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "grpc")
	os.Setenv("AWS_LAMBDA_FUNCTION_NAME", "dummy::local::lambda")

	exporter, err := otlptrace.New(context.Background(), otlptracegrpc.NewClient())
	if err != nil {
		log.Fatalf("failed to initialize stdouttrace export pipeline: %v", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("stdin-project"))),
	)

	otel.SetTracerProvider(tp)

	defer tp.Shutdown(context.Background())

	tr := otel.Tracer("service")
	SetTrancer(tr)

	h := NewBuilder[I, O](srv).
		WithMiddlewares(
			Logger[I, O],
			TraceServiceInvoque[I, O],
		).
		WithMiddlewares(middleware...).
		WithHandlerDecorator(
			OTELDecorator(
				otellambda.WithTracerProvider(tp),
			),
		).
		WithBatch().
		Build()

	http.HandleFunc("/lambda",
		func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			switch req.Method {
			case http.MethodGet:
				writer.WriteHeader(200)
				writer.Write([]byte(`{"status":"running..."}`))
				return
			case http.MethodPost:
				trx, span := Tracer().Start(req.Context(), "service.bootstrap")
				defer span.End()
				raw, err := io.ReadAll(req.Body)
				if err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				ltx := lambdacontext.NewContext(
					trx,
					&lambdacontext.LambdaContext{
						AwsRequestID:       "awsRequestId1234",
						InvokedFunctionArn: "arn:aws:lambda:xxx",
						Identity:           lambdacontext.CognitoIdentity{},
						ClientContext:      lambdacontext.ClientContext{},
					},
				)

				res, err := h.(func(ctx context.Context, in any) (any, error))(ltx, json.RawMessage(raw))
				if err != nil {
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				if err := json.NewEncoder(writer).Encode(res); err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					writer.Write([]byte(fmt.Sprintf(`{"error":%q}`, err.Error())))
					return
				}

				return
			}
		},
	)

	println("listen on port 9090")

	http.ListenAndServe(":9090", nil)
}
