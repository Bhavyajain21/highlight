// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type BackendErrorObjectInput struct {
	SessionSecureID *string       `json:"session_secure_id"`
	RequestID       *string       `json:"request_id"`
	TraceID         *string       `json:"trace_id"`
	SpanID          *string       `json:"span_id"`
	LogCursor       *string       `json:"log_cursor"`
	Event           string        `json:"event"`
	Type            string        `json:"type"`
	URL             string        `json:"url"`
	Source          string        `json:"source"`
	StackTrace      string        `json:"stackTrace"`
	Timestamp       time.Time     `json:"timestamp"`
	Payload         *string       `json:"payload"`
	Service         *ServiceInput `json:"service"`
	Environment     string        `json:"environment"`
}

type ErrorObjectInput struct {
	Event        string             `json:"event"`
	Type         string             `json:"type"`
	URL          string             `json:"url"`
	Source       string             `json:"source"`
	LineNumber   int                `json:"lineNumber"`
	ColumnNumber int                `json:"columnNumber"`
	StackTrace   []*StackFrameInput `json:"stackTrace"`
	Timestamp    time.Time          `json:"timestamp"`
	Payload      *string            `json:"payload"`
}

type InitializeSessionResponse struct {
	SecureID  string `json:"secure_id"`
	ProjectID int    `json:"project_id"`
}

type MetricInput struct {
	SessionSecureID string       `json:"session_secure_id"`
	SpanID          *string      `json:"span_id"`
	ParentSpanID    *string      `json:"parent_span_id"`
	TraceID         *string      `json:"trace_id"`
	Group           *string      `json:"group"`
	Name            string       `json:"name"`
	Value           float64      `json:"value"`
	Category        *string      `json:"category"`
	Timestamp       time.Time    `json:"timestamp"`
	Tags            []*MetricTag `json:"tags"`
}

type MetricTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ReplayEventInput struct {
	Type      int         `json:"type"`
	Timestamp float64     `json:"timestamp"`
	Sid       float64     `json:"_sid"`
	Data      interface{} `json:"data"`
}

type ReplayEventsInput struct {
	Events []*ReplayEventInput `json:"events"`
}

type ServiceInput struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type StackFrameInput struct {
	FunctionName *string       `json:"functionName"`
	Args         []interface{} `json:"args"`
	FileName     *string       `json:"fileName"`
	LineNumber   *int          `json:"lineNumber"`
	ColumnNumber *int          `json:"columnNumber"`
	IsEval       *bool         `json:"isEval"`
	IsNative     *bool         `json:"isNative"`
	Source       *string       `json:"source"`
}

type PublicGraphError string

const (
	PublicGraphErrorBillingQuotaExceeded PublicGraphError = "BillingQuotaExceeded"
)

var AllPublicGraphError = []PublicGraphError{
	PublicGraphErrorBillingQuotaExceeded,
}

func (e PublicGraphError) IsValid() bool {
	switch e {
	case PublicGraphErrorBillingQuotaExceeded:
		return true
	}
	return false
}

func (e PublicGraphError) String() string {
	return string(e)
}

func (e *PublicGraphError) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PublicGraphError(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PublicGraphError", str)
	}
	return nil
}

func (e PublicGraphError) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
