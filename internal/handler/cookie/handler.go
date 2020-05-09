package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	cookie "github.com/neoxelox/microservice-template/internal/entity/cookie"
)

type Handler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Delete(c echo.Context) error
	Modify(c echo.Context) error
}

type CookieHandler struct {
	methods cookie.Methods
}

func NewCookieHandler(methods cookie.Methods) Handler {
	return &CookieHandler{
		methods: methods,
	}
}

func (h *CookieHandler) Create(c echo.Context) error {
	return c.String(http.StatusOK, "Create endpoint!")
}

func (h *CookieHandler) List(c echo.Context) error {
	return c.String(http.StatusOK, "List endpoint!")
}

func (h *CookieHandler) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Delete endpoint!")
}

func (h *CookieHandler) Modify(c echo.Context) error {
	return c.String(http.StatusOK, "Modify endpoint!")
}
