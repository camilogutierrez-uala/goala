package usrv

import (
	"context"
	"encoding/json"
	"errors"
)

type (
	EventProcessor[I any, O any] interface {
		Use(srv Service[I, O])
		Process(ctx context.Context, isBatch bool) (any, error)
	}

	Event[I any, O any] struct {
		event EventProcessor[I, O]
	}
)

func (e *Event[I, O]) Use(srv Service[I, O]) {
	e.event.Use(srv)
}

func (e *Event[I, O]) Process(ctx context.Context, isBatch bool) (any, error) {
	return e.event.Process(ctx, isBatch)
}

func (e *Event[I, O]) UnmarshalJSON(data []byte) error {
	var multi eventMultiMessage[I, O]
	if multi.EventSelector(data) {
		e.event = &multi
		return json.Unmarshal(data, e.event)
	}

	var single eventSingleMessage[I, O]
	if single.EventSelector(data) {
		e.event = &single
		return json.Unmarshal(data, e.event)
	}

	return errors.New("event cannot select")
}
