package usrv

import (
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
)

type (
	EventMultiMessage[I any] interface {
		CanBatch() bool
		Len() int
		Adapt(index int) (*I, error)
		Failure(index int)
		Result() any
	}

	eventMultiMessage[I any, O any] struct {
		event EventMultiMessage[I]
		srv   Service[I, O]
	}
)

const (
	_ddbEvent = "aws:dynamodb"
	_sqsEvent = "aws:sqs"
	_s3Event  = "aws:s3"
	_snsEvent = "aws:sns"
)

var eventKeys = []string{
	"Records.0.eventSource",
	"Records.0.EventSource",
}

func (e *eventMultiMessage[I, O]) EventSelector(raw json.RawMessage) bool {

	for _, key := range eventKeys {
		if source := gjson.GetBytes(raw, key); source.Exists() {
			switch source.String() {
			case _sqsEvent:
				e.event = &EventSQS[I]{}
				return true
			case _ddbEvent:
				e.event = &EventDynamoDB[I]{}
				return true
			case _snsEvent:
				e.event = &EventSNS[I]{}
				return true
			default:
				return false
			}
		}
	}
	return false
}

func (e *eventMultiMessage[I, O]) Process(ctx context.Context, isBatch bool) (any, error) {
	if isBatch && e.event.CanBatch() {
		return e.batch(ctx)
	}

	return nil, e.process(ctx)
}

func (e *eventMultiMessage[I, O]) process(ctx context.Context) error {
	for i := 0; i < e.event.Len(); i++ {
		in, err := e.event.Adapt(i)
		if err != nil {
			return err
		}

		if _, err := e.srv(ctx, in); err != nil {
			return err
		}
	}
	return nil
}

func (e *eventMultiMessage[I, O]) batch(ctx context.Context) (any, error) {
	for i := 0; i < e.event.Len(); i++ {
		in, err := e.event.Adapt(i)
		if err != nil {
			e.event.Failure(i)
			continue
		}

		if _, err := e.srv(ctx, in); err != nil {
			e.event.Failure(i)
			continue
		}
	}
	return e.event.Result(), nil
}

func (e *eventMultiMessage[I, O]) Use(srv Service[I, O]) {
	e.srv = srv
}

func (e *eventMultiMessage[I, O]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, e.event)
}
