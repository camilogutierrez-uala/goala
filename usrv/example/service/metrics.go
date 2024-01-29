package service

import (
	"context"
	"github.com/Bancar/uala-bis-go-dependencies/v2/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/camilogutierrez-uala/goala/ulog"
	"github.com/camilogutierrez-uala/goala/usrv"
	"github.com/camilogutierrez-uala/goala/usrv/metrics"
)

func Metrics() []usrv.Middleware[Request, Response] {
	sess := session.GetSession()
	cw := cloudwatch.New(sess)
	return metrics.NewMeter[Request, Response](
		func(ctx context.Context, in *Request, out *Response, err error) {
			if _, err := cw.PutMetricData(
				&cloudwatch.PutMetricDataInput{
					MetricData: []*cloudwatch.MetricDatum{
						{
							MetricName: aws.String(in.Message),
							Value:      aws.Float64(1),
						},
					},
					Namespace: aws.String("PaymentMethodMetric"),
				},
			); err != nil {
				ulog.Context(ctx).
					With("error", err.Error()).
					Debug("cannot emit metrics")
			}
		},
		func(ctx context.Context, in *Request, out *Response, err error) {
			if _, err := cw.PutMetricData(
				&cloudwatch.PutMetricDataInput{
					MetricData: []*cloudwatch.MetricDatum{
						{
							MetricName: aws.String(out.Status),
							Value:      aws.Float64(1),
						},
					},
					Namespace: aws.String("PaymentMethodMetric"),
				},
			); err != nil {
				ulog.Context(ctx).
					With("error", err.Error()).
					Debug("cannot emit metrics")
			}
		},
	)
}
