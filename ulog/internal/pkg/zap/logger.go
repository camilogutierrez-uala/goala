package zap

import (
	"context"
	"github.com/camilogutierrez-uala/goala/ulog/logger"
	"go.uber.org/zap/zapcore"
	"io"
	"time"

	"go.uber.org/zap"
)

type Zap struct {
	log       *zap.SugaredLogger
	ctxSetter func(context.Context, logger.Logger) logger.Logger
}

func (l *Zap) Context(ctx context.Context) logger.Logger {
	if l.ctxSetter == nil {
		return l
	}

	return l.ctxSetter(ctx, l)
}

func (l *Zap) With(key string, value any) logger.Logger {
	return &Zap{
		log:       l.log.With(zap.Any(key, value)),
		ctxSetter: l.ctxSetter,
	}
}

func (l *Zap) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *Zap) Info(msg string) {
	l.log.Info(msg)
}

func (l *Zap) Error(msg string) {
	l.log.Error(msg)
}

func (l *Zap) Fatal(msg string) {
	l.log.Fatal(msg)
}

func NewZapLogger(ctxSetter func(ctx context.Context, l logger.Logger) logger.Logger, writer io.Writer, level logger.Level) *Zap {
	var (
		encoderCfg = zapcore.EncoderConfig{
			MessageKey:  "data",
			LevelKey:    "logtype",
			TimeKey:     "date",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(time.RFC3339))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
		}
		lvl = func(in logger.Level) zapcore.Level {
			switch in {
			case logger.LevelDebug:
				return zapcore.DebugLevel
			case logger.LevelInfo:
				return zapcore.InfoLevel
			case logger.LevelError:
				return zapcore.ErrorLevel
			case logger.LevelFatal:
				return zapcore.FatalLevel
			default:
				return zapcore.InfoLevel
			}
		}
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(writer),
			lvl(level),
		)
		log = zap.New(core).Sugar()
	)
	return WrapZapLogger(log, ctxSetter)
}

func WrapZapLogger(
	log *zap.SugaredLogger,
	ctxSetter func(ctx context.Context, l logger.Logger) logger.Logger,
) *Zap {
	return &Zap{
		log:       log,
		ctxSetter: ctxSetter,
	}
}
