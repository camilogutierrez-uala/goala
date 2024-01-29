package usrv

import (
	"context"
)

type Meter[I any, O any] func(ctx context.Context, in *I, out *O, err error)

func NewMeter[I any, O any](fn ...Meter[I, O]) (list []Middleware[I, O]) {
	for _, f := range fn {
		list = append(list,
			func(next Service[I, O]) Service[I, O] {
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
