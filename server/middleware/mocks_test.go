package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
	nr "github.com/newrelic/go-agent"
)

type (
	MockLogger struct{}

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

func (m *MockLogger) Output() io.Writer {
	return new(bytes.Buffer)
}

func (m *MockLogger) SetOutput(w io.Writer) {
}

func (m *MockLogger) Prefix() string {
	return ""
}

func (m *MockLogger) SetPrefix(p string) {
}

func (m *MockLogger) Level() log.Lvl {
	return log.INFO
}

func (m *MockLogger) SetLevel(v log.Lvl) {
}

func (m *MockLogger) SetHeader(h string) {
}

func (m *MockLogger) Print(i ...interface{}) {
}

func (m *MockLogger) Printf(format string, args ...interface{}) {

}
func (m *MockLogger) Printj(j log.JSON) {

}
func (m *MockLogger) Debug(i ...interface{}) {

}
func (m *MockLogger) Debugf(format string, args ...interface{}) {
}

func (m *MockLogger) Debugj(j log.JSON) {
}

func (m *MockLogger) Info(i ...interface{}) {
}

func (m *MockLogger) Infof(format string, args ...interface{}) {
}

func (m *MockLogger) Infoj(j log.JSON) {
}

func (m *MockLogger) Warn(i ...interface{}) {
}

func (m *MockLogger) Warnf(format string, args ...interface{}) {
}

func (m *MockLogger) Warnj(j log.JSON) {
}

func (m *MockLogger) Error(i ...interface{}) {
}

func (m *MockLogger) Errorf(format string, args ...interface{}) {
}

func (m *MockLogger) Errorj(j log.JSON) {
}

func (m *MockLogger) Fatal(i ...interface{}) {
}

func (m *MockLogger) Fatalj(j log.JSON) {
}

func (m *MockLogger) Fatalf(format string, args ...interface{}) {
}

func (m *MockLogger) Panic(i ...interface{}) {
}

func (m *MockLogger) Panicj(j log.JSON) {
}

func (m *MockLogger) Panicf(format string, args ...interface{}) {
}
