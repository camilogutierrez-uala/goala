package ulog

import (
	"bytes"
	"context"
	"testing"
)

func TestMethods(t *testing.T) {
	logBuffer := &bytes.Buffer{}
	SetWriter(logBuffer)
	SetLevel(1)

	cases := []struct {
		name string
		call func()
		out  string
	}{
		{
			name: "Debug call",
			call: func() {
				Debug("testDebug")
			},
			out: "Debug: testDebug\n",
		},
		{
			name: "Info call",
			call: func() {
				Info("testInfo")
			},
			out: "Info: testInfo\n",
		},
		{
			name: "Error call",
			call: func() {
				Error("testError")
			},
			out: "Error: testError\n",
		},
		{
			name: "Fatal call",
			call: func() {
				Fatal("testFatal")
			},
			out: "Fatal: testFatal\n",
		},
		{
			name: "With call",
			call: func() {
				With("key", "value").Debug("testWith")
			},
			out: "Debug: testWith\n",
		},
		{
			name: "Context call",
			call: func() {
				Context(context.Background()).Debug("testContext")
			},
			out: "Debug: testContext\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			logBuffer.Reset()
			c.call()
			if got := logBuffer.String(); got != c.out {
				t.Errorf("Expected %q but got %q", c.out, got)
			}
		})
	}
}
