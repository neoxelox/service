package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

func NewMockHandler() *MockHandler {
	return &MockHandler{}
}

func (h *MockHandler) Create(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}

func (h *MockHandler) List(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}

func (h *MockHandler) Delete(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}

func (h *MockHandler) Modify(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}
