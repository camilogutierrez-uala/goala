package logger

import (
	"context"
)

const (
	LevelDebug Level = iota + 1
	LevelInfo
	LevelError
	LevelFatal
)

type (
	Level int

	Logger interface {
		Context(ctx context.Context) Logger
		With(key string, value any) Logger
		Debug(msg string)
		Info(msg string)
		Error(msg string)
		Fatal(msg string)
	}
)
