package ulog

import (
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/camilogutierrez-uala/goala/ulog/logger"
)

var contextAdapter = func(ctx context.Context, l logger.Logger) logger.Logger {
	trx, ok := lambdacontext.FromContext(ctx)
	if ok {
		return l.
			With("AWS_REQUEST_ID", trx.AwsRequestID).
			With("FUNCTION_ARN", trx.InvokedFunctionArn)
	}
	return l
}

func SetContextAdapter(adapt func(ctx context.Context, l logger.Logger) logger.Logger) {
	contextAdapter = adapt
}

func UseContext(ctx context.Context, l logger.Logger) logger.Logger {
	return contextAdapter(ctx, l)
}
