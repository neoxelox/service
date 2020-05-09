// From https://github.com/dafiti/echo-middleware
// MIT License

package middleware

import (
	"bytes"
	"io"

	"github.com/labstack/gommon/log"
)

type MockLogger struct{}

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
