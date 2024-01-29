package usrv

import (
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
)

type (
	EventSingleMessage[I any] interface {
		Adapt() (*I, error)
	}

	eventSingleMessage[I any, O any] struct {
		event EventSingleMessage[I]
		srv   Service[I, O]
	}
)

func (e *eventSingleMessage[I, O]) EventSelector(raw json.RawMessage) bool {
	if source := gjson.GetBytes(raw, "source"); source.Exists() {
		e.event = &EventCloudWatch[I]{}
		return true
	}

	e.event = &EventJSON[I]{}
	return true
}

func (e *eventSingleMessage[I, O]) Process(ctx context.Context, _ bool) (any, error) {
	in, err := e.event.Adapt()
	if err != nil {
		return nil, err
	}

	return e.srv(ctx, in)
}

func (e *eventSingleMessage[I, O]) Use(srv Service[I, O]) {
	e.srv = srv
}

func (e *eventSingleMessage[I, O]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, e.event)
}
