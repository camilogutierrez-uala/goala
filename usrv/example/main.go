package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/camilogutierrez-uala/goala/usrv"
	"github.com/camilogutierrez-uala/goala/usrv/example/service"
	"github.com/camilogutierrez-uala/goala/usrv/otel"
)

func main() {
	HTTPLocalLambda()
}

func EventSQS() {
	raw := `
{
  "Records": [
    {
      "messageId": "19dd0b57-b21e-4ac1-bd88-01bbb068cb78",
      "receiptHandle": "MessageReceiptHandle",
      "body": "\n{\n  \"Records\": [\n    {\n      \"messageId\": \"19dd0b57-b21e-4ac1-bd88-01bbb068cb78\",\n      \"receiptHandle\": \"MessageReceiptHandle\",\n      \"body\": \"{\\\"foo\\\":\\\"foo\\\"}\",\n      \"attributes\": {\n        \"ApproximateReceiveCount\": \"1\",\n        \"SentTimestamp\": \"1523232000000\",\n        \"SenderId\": \"123456789012\",\n        \"ApproximateFirstReceiveTimestamp\": \"1523232000001\"\n      },\n      \"messageAttributes\": {},\n      \"md5OfBody\": \"7b270e59b47ff90a553787216d55d91d\",\n      \"eventSource\": \"aws:sqs\",\n      \"eventSourceARN\": \"arn:{partition}:sqs:{region}:123456789012:MyQueue\",\n      \"awsRegion\": \"{region}\"\n    }\n  ]\n}\n",
      "attributes": {
        "ApproximateReceiveCount": "1",
        "SentTimestamp": "1523232000000",
        "SenderId": "123456789012",
        "ApproximateFirstReceiveTimestamp": "1523232000001"
      },
      "messageAttributes": {},
      "md5OfBody": "7b270e59b47ff90a553787216d55d91d",
      "eventSource": "aws:sqs",
      "eventSourceARN": "arn:{partition}:sqs:{region}:123456789012:MyQueue",
      "awsRegion": "{region}"
    }
  ]
}
`
	var event usrv.EventSQS[usrv.EventSQS[any]]
	if err := json.Unmarshal([]byte(raw), &event); err != nil {
		panic(err)
	}

	in, err := event.Adapt(0)
	if err != nil {
		panic(err)
	}

	on, err := in.Adapt(0)
	if err != nil {
		panic(err)
	}

	println(raw, in, on)
}

func EventDynamoDB() {
	raw := `
{
  "Records": [
    {
      "eventID": "1",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    },
    {
      "eventID": "1",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    },
    {
      "eventID": "1",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    },
    {
      "eventID": "1",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    }
  ]
}
`

	type DynamoDBItem struct {
		Number    float64        `dynamodbav:"Number"`
		String    string         `dynamodbav:"String"`
		Binary    []byte         `dynamodbav:"Binary"`
		Boolean   bool           `dynamodbav:"Boolean"`
		Null      any            `dynamodbav:"Null"`
		List      []any          `dynamodbav:"List"`
		Map       map[string]any `dynamodbav:"Map"`
		NumberSet []float64      `dynamodbav:"NumberSet"`
		StringSet []string       `dynamodbav:"StringSet"`
		BinarySet [][]byte       `dynamodbav:"BinarySet"`
	}

	var event usrv.EventDynamoDB[DynamoDBItem]
	if err := json.Unmarshal([]byte(raw), &event); err != nil {
		panic(err)
	}

	in, err := event.Adapt(0)
	if err != nil {
		panic(err)
	}

	println(raw, in)
}

func EventAgnostic() {
	raw := `
{
  "Records": [
    {
      "eventID": "1",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    },
    {
      "eventID": "2",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    },
    {
      "eventID": "3",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, World!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    },
    {
      "eventID": "4",
      "eventName": "INSERT",
      "eventVersion": "1.0",
      "eventSource": "aws:dynamodb",
      "awsRegion": "{region}",
      "dynamodb": {
        "Keys": {
          "Id": {
            "N": "101"
          }
        },
        "NewImage": {
          "Number": {
            "N": "123.456"
          },
          "String": {
            "S": "Hello, fail!"
          },
          "Binary": {
            "B": "dGhpcyBpcyBhIHRlc3Q="
          },
          "Boolean": {
            "BOOL": true
          },
          "Null": {
            "NULL": true
          },
          "List": {
            "L": [
              {
                "N": "123"
              },
              {
                "S": "foobar"
              },
              { 
                "M": {
                  "foo": {
                    "BOOL": true
                  },
                  "bar": {
                    "N": "789"
                  },
                  "baz": {
                    "S": "qux"
                  }
                }
              }
            ]
          },
          "Map": {
            "M": {
              "foo": {
                "BOOL": true
              },
              "bar": {
                "N": "789"
              },
              "baz": {
                "S": "qux"
              }
            }
          },
          "NumberSet": {
            "NS": [
              "42",
              "123"
            ]
          },
          "StringSet": {
            "SS": [
              "Hello",
              "World"
            ]
          },
          "BinarySet": {
            "BS": [
              "QmFzZTY0IGVuY29kZWQgc3RyaW5n"
            ]
          }
        },
        "SequenceNumber": "111",
        "SizeBytes": 26,
        "StreamViewType": "NEW_AND_OLD_IMAGES"
      },
      "eventSourceARN": "arn:{partition}:dynamodb:{region}:account-id:table/ExampleTableWithStream/stream/2015-06-27T00:48:05.899"
    }
  ]
}
`

	type DynamoDBItem struct {
		Number    float64        `dynamodbav:"Number"`
		String    string         `dynamodbav:"String"`
		Binary    []byte         `dynamodbav:"Binary"`
		Boolean   bool           `dynamodbav:"Boolean"`
		Null      any            `dynamodbav:"Null"`
		List      []any          `dynamodbav:"List"`
		Map       map[string]any `dynamodbav:"Map"`
		NumberSet []float64      `dynamodbav:"NumberSet"`
		StringSet []string       `dynamodbav:"StringSet"`
		BinarySet [][]byte       `dynamodbav:"BinarySet"`
	}

	var event usrv.Event[DynamoDBItem, any]
	if err := json.Unmarshal([]byte(raw), &event); err != nil {
		panic(err)
	}

	event.Use(
		func(ctx context.Context, in *DynamoDBItem) (*any, error) {
			var asd any = "dsadsadasd"
			if in.String == "Hello, fail!" {
				return nil, errors.New("error")
			}
			return &asd, nil
		},
	)

	in, err := event.Process(context.Background(), true)
	if err != nil {
		panic(err)
	}

	println(raw, in)
}

func HTTPLocalLambda() {
	type DynamoDBItem struct {
		Number    float64        `dynamodbav:"Number"`
		String    string         `dynamodbav:"String"`
		Binary    []byte         `dynamodbav:"Binary"`
		Boolean   bool           `dynamodbav:"Boolean"`
		Null      any            `dynamodbav:"Null"`
		List      []any          `dynamodbav:"List"`
		Map       map[string]any `dynamodbav:"Map"`
		NumberSet []float64      `dynamodbav:"NumberSet"`
		StringSet []string       `dynamodbav:"StringSet"`
		BinarySet [][]byte       `dynamodbav:"BinarySet"`
	}

	srv := func(ctx context.Context, in *DynamoDBItem) (*any, error) {
		var asd any = "dsadsadasd"
		if in.String == "Hello, fail!" {
			return nil, errors.New("error")
		}
		return &asd, nil
	}

	usrv.LocalHTTP(srv, true)
}

func BaseLambda() {
	srv := service.New()
	usrv.LambdaServe(srv.Service, service.Metrics()...)
}

func OtelLambda() {
	srv := service.New()
	if err := otel.UseTrace(); err != nil {
		panic(err)
		return
	}
	defer otel.Shutdown()

	usrv.LambdaOTELServe(srv.Service, otel.LambdaOptions(), service.Metrics()...)
}
