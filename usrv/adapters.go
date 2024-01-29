package usrv

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type (
	AdapterFn[I any] func(raw json.RawMessage) (*I, error)

	EventJSON[I any] struct {
		json.RawMessage
	}

	EventCloudWatch[I any] struct {
		*events.CloudWatchEvent
	}

	EventSQS[I any] struct {
		*events.SQSEvent
		failures []events.SQSBatchItemFailure
	}

	EventDynamoDB[I any] struct {
		*events.DynamoDBEvent
		failures []events.DynamoDBBatchItemFailure
	}

	EventSNS[I any] struct {
		*events.SNSEvent
	}

	EventS3[I any] struct {
		*events.S3Event
	}
)

func (e *EventSNS[I]) CanBatch() bool {
	return false
}

func (e *EventSNS[I]) Len() int {
	return len(e.Records)
}

func (e *EventSNS[I]) Adapt(index int) (*I, error) {
	record := e.Records[index]
	raw := bytes.NewBufferString(record.SNS.Message).Bytes()
	return AdaptJSON[I](raw)
}

func (e *EventSNS[I]) Failure(_ int) {
}

func (e *EventSNS[I]) Result() any {
	return nil
}

func (e *EventDynamoDB[I]) CanBatch() bool {
	return true
}

func (e *EventDynamoDB[I]) Len() int {
	return len(e.Records)
}

func (e *EventDynamoDB[I]) Adapt(index int) (*I, error) {
	record := e.Records[index]
	attrs := make(map[string]types.AttributeValue, 0)
	for k, v := range record.Change.NewImage {
		attrs[k] = e.ToTypes(v)
	}

	var in I
	if err := attributevalue.UnmarshalMap(attrs, &in); err != nil {
		return nil, err
	}

	return &in, nil
}

func (e *EventDynamoDB[I]) ToTypes(record events.DynamoDBAttributeValue) types.AttributeValue {
	var val types.AttributeValue

	switch record.DataType() {
	case events.DataTypeBinary:
		val = &types.AttributeValueMemberB{
			Value: record.Binary(),
		}
	case events.DataTypeBinarySet:
		val = &types.AttributeValueMemberBS{
			Value: record.BinarySet(),
		}
	case events.DataTypeBoolean:
		val = &types.AttributeValueMemberBOOL{
			Value: record.Boolean(),
		}
	case events.DataTypeList:
		var items []types.AttributeValue
		for _, value := range record.List() {
			items = append(items, e.ToTypes(value))
		}
		val = &types.AttributeValueMemberL{
			Value: items,
		}
	case events.DataTypeMap:
		items := make(map[string]types.AttributeValue)
		for k, v := range record.Map() {
			items[k] = e.ToTypes(v)
		}
		val = &types.AttributeValueMemberM{
			Value: items,
		}
	case events.DataTypeNull:
		val = nil
	case events.DataTypeNumber:
		val = &types.AttributeValueMemberN{
			Value: record.Number(),
		}
	case events.DataTypeNumberSet:
		val = &types.AttributeValueMemberNS{
			Value: record.NumberSet(),
		}
	case events.DataTypeString:
		val = &types.AttributeValueMemberS{
			Value: record.String(),
		}
	case events.DataTypeStringSet:
		val = &types.AttributeValueMemberSS{
			Value: record.StringSet(),
		}
	}

	return val
}

func (e *EventDynamoDB[I]) Failure(index int) {
	e.failures = append(
		e.failures,
		events.DynamoDBBatchItemFailure{
			ItemIdentifier: e.Records[index].EventID,
		},
	)
}

func (e *EventDynamoDB[I]) Result() any {
	return e.failures
}

func (e *EventSQS[I]) CanBatch() bool {
	return true
}

func (e *EventSQS[I]) Len() int {
	return len(e.Records)
}

func (e *EventSQS[I]) Adapt(index int) (*I, error) {
	record := e.Records[index]
	raw := bytes.NewBufferString(record.Body).Bytes()
	return AdaptJSON[I](raw)
}

func (e *EventSQS[I]) Failure(index int) {
	e.failures = append(
		e.failures,
		events.SQSBatchItemFailure{
			ItemIdentifier: e.Records[index].MessageId,
		},
	)
}

func (e *EventSQS[I]) Result() any {
	return &events.SQSEventResponse{
		BatchItemFailures: e.failures,
	}
}

func (e *EventCloudWatch[I]) Adapt() (*I, error) {
	return AdaptJSON[I](e.Detail)
}

func (e *EventJSON[I]) Adapt() (*I, error) {
	return AdaptJSON[I](e.RawMessage)
}

func AdaptJSON[I any](raw json.RawMessage) (*I, error) {
	var in I
	if err := json.Unmarshal(raw, &in); err != nil {
		return nil, err
	}
	return &in, nil
}
