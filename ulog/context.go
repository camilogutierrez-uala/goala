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
			With("request.awsRequestId", trx.AwsRequestID).
			With("function.arn", trx.InvokedFunctionArn)
	}
	return l
}

func SetContextAdapter(adapt func(ctx context.Context, l logger.Logger) logger.Logger) {
	contextAdapter = adapt
}

func UseContext(ctx context.Context, l logger.Logger) logger.Logger {
	return contextAdapter(ctx, l)
}
