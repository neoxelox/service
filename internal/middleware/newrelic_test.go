// From https://github.com/dafiti/echo-middleware
// MIT License

package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewRelicWhenRaiseAError(t *testing.T) {
	defer func() {
		err := recover().(error)

		if err.Error() != "New relic: license length is not 40" {
			t.Fatalf("Wrong panic error: %s", err.Error())
		}
	}()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewRelic("ok", "ok")(func(c echo.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	})

	h(c)
}

func TestNewRelicWithApplication(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/something", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	app := new(Application)

	h := NewRelicWithApplication(app)(func(c echo.Context) error {
		return fmt.Errorf("Something wrong")
	})

	h(c)
}

func TestNewrelicWithHttpMethod(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/test")

	app := new(Application)

	h := NewRelicWithApplication(app)(func(c echo.Context) error {
		txn := c.Get(NEWRELIC_TXN).(*Transaction)

		if txn.Name != "/test [GET]" {
			t.Fatalf("Invalid transaction name: %s", txn.Name)
		}

		return fmt.Errorf("Something wrong")
	})

	h(c)
}
