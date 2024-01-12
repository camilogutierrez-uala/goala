package usrv

import (
	"context"
	"github.com/camilogutierrez-uala/goala/ulog"
)

func Logger[I any, O any](next HandlerFn[I, O]) HandlerFn[I, O] {
	return func(ctx context.Context, in *I) (out *O, err error) {
		defer func() {
			log := ulog.With(
				zap.Any("event_handler_input", in),
			)
			if err != nil {
				log.With(
					"error", err,
				).Error("event handling has an error")
			} else {
				log.With(
					zap.Any("event_handler_output", out),
				).Info("event handling success")
			}
		}()
		return next(ctx, in)
	}
}
