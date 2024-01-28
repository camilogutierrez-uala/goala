package ulog

import (
	"bytes"
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/camilogutierrez-uala/goala/ulog/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestMethods is a unit test function for testing the logging methods of a logger.
//
// It sets the environment variables required for logging and initializes the logger.
// Then, it runs a series of test cases where each test logs a message using a different logging method.
// The function checks if the logged message matches the expected result using regular expressions.
//
// Test cases:
// - Context log: Logs a message using the Context method and checks if the logged message matches the expected format.
// - Debug log: Logs a message using the Debug method and checks if the logged message matches the expected format.
// - Error log: Logs a message using the Error method and checks if the logged message matches the expected format.
// - Info log: Logs a message using the Info method and checks if the logged message matches the expected format.
// - With log: Logs a message using the With method with a key-value pair and checks if the logged message matches the expected format.
func TestMethods(t *testing.T) {

	os.Setenv("AWS_LAMBDA_FUNCTION_NAME", "aws_lambda_function_name")
	os.Setenv("AWS_LAMBDA_FUNCTION_VERSION", "aws_lambda_function_version")
	os.Setenv("ENVIRONMENT", "environment")
	os.Setenv("FLOW", "flow")
	os.Setenv("FUNCTION_ARN", "function_arn")
	os.Setenv("COUNTRY", "country")
	os.Setenv("AWS_REQUEST_ID", "aws_request_id")
	os.Setenv("UALA_REQUEST_ID", "uala_request_id")
	os.Setenv("USERNAME", "username")
	os.Setenv("ACCOUNT_ID", "account_id")

	logBuffer := &bytes.Buffer{}
	SetWriter(logBuffer)
	SetLevel(logger.LevelDebug)

	tests := []struct {
		name string
		log  func()
		want string
	}{
		{
			name: "Context log",
			log: func() {
				ctx := &lambdacontext.LambdaContext{
					AwsRequestID:       "awsRequestId1234",
					InvokedFunctionArn: "arn:aws:lambda:xxx",
					Identity:           lambdacontext.CognitoIdentity{},
					ClientContext:      lambdacontext.ClientContext{},
				}
				Context(lambdacontext.NewContext(context.TODO(), ctx)).Debug("testContext")
			},
			want: `{"time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,6}-\d{2}:\d{2}","level":"DEBUG","msg":"testContext","function.name":"aws_lambda_function_name","function.version":"aws_lambda_function_version","function.env":"environment","function.flow":"flow","function.arn":"function_arn","function.country":"country","request.awsRequestId":"aws_request_id","request.ualaRequestId":"uala_request_id","user.username":"username","user.accountId":"account_id","request.awsRequestId":"awsRequestId1234","function.arn":"arn:aws:lambda:xxx"}` + "\n",
		},
		{
			name: "Debug log",
			log: func() {
				Debug("testDebug")
			},
			want: `{"time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,6}-\d{2}:\d{2}","level":"DEBUG","msg":"testDebug","function.name":"aws_lambda_function_name","function.version":"aws_lambda_function_version","function.env":"environment","function.flow":"flow","function.arn":"function_arn","function.country":"country","request.awsRequestId":"aws_request_id","request.ualaRequestId":"uala_request_id","user.username":"username","user.accountId":"account_id"}` + "\n",
		},
		{
			name: "Error log",
			log: func() {
				Error("testError")
			},
			want: `{"time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,6}-\d{2}:\d{2}","level":"ERROR","msg":"testError","function.name":"aws_lambda_function_name","function.version":"aws_lambda_function_version","function.env":"environment","function.flow":"flow","function.arn":"function_arn","function.country":"country","request.awsRequestId":"aws_request_id","request.ualaRequestId":"uala_request_id","user.username":"username","user.accountId":"account_id"}` + "\n",
		},
		{
			name: "Info log",
			log: func() {
				Info("testInfo")
			},
			want: `{"time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,6}-\d{2}:\d{2}","level":"INFO","msg":"testInfo","function.name":"aws_lambda_function_name","function.version":"aws_lambda_function_version","function.env":"environment","function.flow":"flow","function.arn":"function_arn","function.country":"country","request.awsRequestId":"aws_request_id","request.ualaRequestId":"uala_request_id","user.username":"username","user.accountId":"account_id"}` + "\n",
		},
		{
			name: "With log",
			log: func() {
				With("key", "value").Debug("testWith")
			},
			want: `{"time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,6}-\d{2}:\d{2}","level":"DEBUG","msg":"testWith","function.name":"aws_lambda_function_name","function.version":"aws_lambda_function_version","function.env":"environment","function.flow":"flow","function.arn":"function_arn","function.country":"country","request.awsRequestId":"aws_request_id","request.ualaRequestId":"uala_request_id","user.username":"username","user.accountId":"account_id","key":"value"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logBuffer.Reset()
			tt.log()
			got := logBuffer.String()
			assert.Regexp(t, tt.want, got)
		})
	}
}
