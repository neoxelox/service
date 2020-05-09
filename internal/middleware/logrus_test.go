// From https://github.com/dafiti/echo-middleware
// MIT License

package middleware

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func TestLogrusWithConfig(t *testing.T) {
	e := echo.New()

	form := url.Values{}
	form.Add("username", "doejohn")

	req := httptest.NewRequest(echo.POST, "http://some?name=john", strings.NewReader(form.Encode()))

	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Header.Add(echo.HeaderXRequestID, "123")
	req.Header.Add("Referer", "http://foo.bar")
	req.Header.Add("User-Agent", "cli-agent")
	req.Header.Add(echo.HeaderXForwardedFor, "http://foo.bar")
	req.Header.Add("store", "dafiti")
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: "A1B2C3",
	})

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	b := new(bytes.Buffer)

	logger := logrus.StandardLogger()
	logger.Out = b

	fields := DefaultLogrusConfig.FieldMap
	fields["empty"] = ""
	fields["path"] = "@path"
	fields["referer"] = "@referer"
	fields["user_agent"] = "@user_agent"
	fields["store"] = "@header:store"
	fields["filter_name"] = "@query:name"
	fields["username"] = "@form:username"
	fields["session"] = "@cookie:session"

	config := LogrusConfig{
		Logger:   logger,
		FieldMap: fields,
	}

	LogrusWithConfig(config)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)

	res := b.String()

	if !strings.Contains(res, "Handle request") {
		t.Errorf("Invalid log response body, handle request info not found")
	}

	if !strings.Contains(res, "id=123") {
		t.Errorf("Invalid log response body, request id not found")
	}

	if !strings.Contains(res, `remote_ip="http://foo.bar"`) {
		t.Errorf("Invalid log response body, remote ip not found")
	}

	if !strings.Contains(res, `uri="http://some?name=john"`) {
		t.Errorf("Invalid log response body, uri not found")
	}

	if !strings.Contains(res, "host=some") {
		t.Errorf("Invalid log response body, host not found")
	}

	if !strings.Contains(res, "method=POST") {
		t.Errorf("Invalid log response body, method not found")
	}

	if !strings.Contains(res, "status=200") {
		t.Errorf("Invalid log response body, status not found")
	}

	if !strings.Contains(res, "latency=") {
		t.Errorf("Invalid log response body, latency not found")
	}

	if !strings.Contains(res, "latency_human=") {
		t.Errorf("Invalid log response body, latency_human not found")
	}

	if !strings.Contains(res, "bytes_in=0") {
		t.Errorf("Invalid log response body, bytes_in not found")
	}

	if !strings.Contains(res, "bytes_out=4") {
		t.Errorf("Invalid log response body, bytes_out not found")
	}

	if !strings.Contains(res, "path=/") {
		t.Errorf("Invalid log response body, path not found")
	}

	if !strings.Contains(res, `referer="http://foo.bar"`) {
		t.Errorf("Invalid log response body, referer not found")
	}

	if !strings.Contains(res, "user_agent=cli-agent") {
		t.Errorf("Invalid log response body, user_agent not found")
	}

	if !strings.Contains(res, "store=dafiti") {
		t.Errorf("Invalid log response body, header store not found")
	}

	if !strings.Contains(res, "filter_name=john") {
		t.Errorf("Invalid log response body, query filter_name not found")
	}

	if !strings.Contains(res, "username=doejohn") {
		t.Errorf("Invalid log response body, form field username not found")
	}

	if !strings.Contains(res, "session=A1B2C3") {
		t.Errorf("Invalid log response body, cookie session not found")
	}
}

func TestLogrus(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	Logrus()(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)
}

func TestLogrusWithEmptyConfig(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	LogrusWithConfig(LogrusConfig{})(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)
}

func TestLogrusRetrievesAnError(t *testing.T) {
	e := echo.New()
	e.Logger = &MockLogger{}
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	b := new(bytes.Buffer)

	logger := logrus.StandardLogger()
	logger.Out = b

	config := LogrusConfig{
		Logger: logger,
	}

	LogrusWithConfig(config)(func(c echo.Context) error {
		return errors.New("error")
	})(c)

	res := b.String()

	if !strings.Contains(res, "status=500") {
		t.Errorf("Invalid log response body, wrong status code")
	}
}

func TestLogrusWithSkipper(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/some", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	config := DefaultLogrusConfig
	config.Skipper = func(c echo.Context) bool {
		return true
	}

	LogrusWithConfig(config)(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})(c)
}
