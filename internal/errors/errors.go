package errors

import "github.com/labstack/echo/v4"

type Error struct {
	err         string
	description string
	level       string
	tags        map[string]string
	autoLog     bool
	autoPanic   bool
}

type New struct {
	Err         string
	Description string
	Level       string
	Tags        map[string]string
	AutoLog     bool
	AutoPanic   bool
}

func (n New) Error() Error {
	return Error{}
}

func (e Error) Error() string {
	return e.err
}

func (e Error) String() string {
	return e.err
}

// TODO: with or without echo context
func (e Error) Raise(context echo.Context, extra map[string]string) error {
	return e
}

func (e Error) Log() {

}
