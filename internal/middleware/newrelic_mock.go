// From https://github.com/dafiti/echo-middleware
// MIT License

package middleware

import (
	"net/http"
	"time"

	nr "github.com/newrelic/go-agent"
)

type (
	Application struct{}

	Transaction struct {
		http.ResponseWriter
		Name string
	}

	DistributedTracePayload struct{}
)

func (app *Application) StartTransaction(name string, w http.ResponseWriter, r *http.Request) nr.Transaction {
	return &Transaction{Name: name}
}

func (app *Application) RecordCustomEvent(eventType string, params map[string]interface{}) error {
	return nil
}

func (app *Application) Shutdown(timeout time.Duration) {
}

func (app *Application) WaitForConnection(timeout time.Duration) error {
	return nil
}

func (app *Application) RecordCustomMetric(name string, value float64) error {
	return nil
}

func (tn *Transaction) Application() nr.Application {
	return &Application{}
}

func (tn *Transaction) End() error {
	return nil
}

func (tn *Transaction) Ignore() error {
	return nil
}

func (tn *Transaction) SetName(name string) error {
	return nil
}

func (tn *Transaction) NoticeError(err error) error {
	return nil
}

func (tn *Transaction) AddAttribute(key string, value interface{}) error {
	return nil
}

func (tn *Transaction) AcceptDistributedTracePayload(t nr.TransportType, payload interface{}) error {
	return nil
}

func (tn *Transaction) StartSegmentNow() nr.SegmentStartTime {
	return nr.SegmentStartTime{}
}

func (tn *Transaction) BrowserTimingHeader() (*nr.BrowserTimingHeader, error) {
	return &nr.BrowserTimingHeader{}, nil
}

func (tn *Transaction) GetLinkingMetadata() nr.LinkingMetadata {
	return nr.LinkingMetadata{}
}

func (tn *Transaction) GetTraceMetadata() nr.TraceMetadata {
	return nr.TraceMetadata{}
}

func (tn *Transaction) IsSampled() bool {
	return false
}

func (tn *Transaction) SetWebRequest(nr.WebRequest) error {
	return nil
}

func (tn *Transaction) CreateDistributedTracePayload() nr.DistributedTracePayload {
	return &DistributedTracePayload{}
}

func (dt *DistributedTracePayload) HTTPSafe() string {
	return ""
}

func (dt *DistributedTracePayload) Text() string {
	return ""
}

func (tn *Transaction) NewGoroutine() nr.Transaction {
	return nil
}

func (tn *Transaction) SetWebResponse(http.ResponseWriter) nr.Transaction {
	return nil
}
