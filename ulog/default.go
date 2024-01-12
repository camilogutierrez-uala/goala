package ulog

import (
	"context"
	"github.com/camilogutierrez-uala/goala/ulog/internal/pkg/zap"
	"github.com/camilogutierrez-uala/goala/ulog/logger"
	"io"
	"os"
)

var (
	_loggerBuilder           = zap.NewZapLogger
	_defaultLevel            = logger.LevelInfo
	_defaultWriter io.Writer = os.Stdout
	_defaultLog              = func() logger.Logger {
		log := _loggerBuilder(UseContext, _defaultWriter, _defaultLevel)
		return SetEnvironment(log)
	}()
)

func SetEnvironment(l logger.Logger) logger.Logger {
	return l.With("function.name", os.Getenv("AWS_LAMBDA_FUNCTION_NAME")).
		With("function.version", os.Getenv("AWS_LAMBDA_FUNCTION_VERSION")).
		With("function.env", os.Getenv("ENVIRONMENT")).
		With("function.flow", os.Getenv("FLOW")).
		With("function.arn", os.Getenv("FUNCTION_ARN")).
		With("function.country", os.Getenv("COUNTRY")).
		With("request.awsRequestId", os.Getenv("AWS_REQUEST_ID")).
		With("request.ualaRequestId", os.Getenv("UALA_REQUEST_ID")).
		With("user.username", os.Getenv("USERNAME")).
		With("user.accountId", os.Getenv("ACCOUNT_ID"))
}

func SetLevel(lvl logger.Level) {
	_defaultLevel = lvl
	_defaultLog = SetEnvironment(
		_loggerBuilder(UseContext, _defaultWriter, _defaultLevel),
	)
}

func SetWriter(writer io.Writer) {
	_defaultWriter = writer
	_defaultLog = SetEnvironment(
		_loggerBuilder(UseContext, _defaultWriter, _defaultLevel),
	)
}

func Context(ctx context.Context) logger.Logger {
	return _defaultLog.Context(ctx)
}

func With(key string, value any) logger.Logger {
	return _defaultLog.With(key, value)
}

func Debug(msg string) {
	_defaultLog.Debug(msg)
}

func Info(msg string) {
	_defaultLog.Info(msg)
}

func Error(msg string) {
	_defaultLog.Error(msg)
}

func Fatal(msg string) {
	_defaultLog.Fatal(msg)
}
