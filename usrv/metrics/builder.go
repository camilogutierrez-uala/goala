package metrics

import (
	"context"
	"github.com/Bancar/goala/usrv"
)

type Meter[I any, O any] func(ctx context.Context, in *I, out *O, err error)

func NewMeter[I any, O any](fn ...Meter[I, O]) (list []usrv.Middleware[I, O]) {
	for _, f := range fn {
		list = append(list,
			func(next usrv.Service[I, O]) usrv.Service[I, O] {
				return func(ctx context.Context, in *I) (out *O, err error) {
					defer func() {
						f(ctx, in, out, err)
					}()
					return next(ctx, in)
				}
			},
		)
	}
	return
}
