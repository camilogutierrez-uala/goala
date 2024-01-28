package slog

import (
	"context"
	"github.com/camilogutierrez-uala/goala/ulog/logger"
	"io"
	"log/slog"
)

type Slog struct {
	log       *slog.Logger
	ctxSetter func(ctx context.Context, logger logger.Logger) logger.Logger
}

func (s *Slog) Context(ctx context.Context) logger.Logger {
	if s.ctxSetter == nil {
		return s
	}
	return s.ctxSetter(ctx, s)
}

func (s *Slog) With(key string, value any) logger.Logger {
	log := s.log.With(key, value)
	return &Slog{
		log:       log,
		ctxSetter: s.ctxSetter,
	}
}

func (s *Slog) Debug(msg string) {
	s.log.Debug(msg)
}

func (s *Slog) Info(msg string) {
	s.log.Info(msg)
}

func (s *Slog) Error(msg string) {
	s.log.Error(msg)
}

func NewSlogLogger(ctxSetter func(ctx context.Context, l logger.Logger) logger.Logger, writer io.Writer, level logger.Level) *Slog {
	var (
		lvl = func(in logger.Level) slog.Level {
			switch in {
			case logger.LevelDebug:
				return slog.LevelDebug
			case logger.LevelInfo:
				return slog.LevelInfo
			case logger.LevelError:
				return slog.LevelError
			default:
				return slog.LevelInfo
			}
		}(level)
		log = slog.New(
			slog.NewJSONHandler(
				writer,
				&slog.HandlerOptions{
					AddSource: false,
					Level:     lvl,
				},
			),
		)
	)
	return wrapLogger(log, ctxSetter)
}

func wrapLogger(
	log *slog.Logger,
	ctxSetter func(ctx context.Context, l logger.Logger) logger.Logger,
) *Slog {
	return &Slog{
		log:       log,
		ctxSetter: ctxSetter,
	}
}
