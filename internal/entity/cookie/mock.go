package entity

import (
	"context"

	"github.com/rs/xid"
	"github.com/stretchr/testify/mock"
)

type MockMethods struct {
	mock.Mock
}

func NewMockMethods() *MockMethods {
	return &MockMethods{}
}

func (m *MockMethods) Create(ctx context.Context, title string,
	category *string) (*Cookie, error) {
	args := m.Called(ctx, title, category)
	return args.Get(0).(*Cookie), args.Error(1)
}

func (m *MockMethods) ListAll(ctx context.Context) ([]*Cookie, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Cookie), args.Error(1)
}

func (m *MockMethods) Delete(ctx context.Context, cookieID xid.ID) error {
	args := m.Called(ctx, cookieID)
	return args.Error(0)
}

func (m *MockMethods) Modify(ctx context.Context, cookieID xid.ID,
	title string, category *string) (*Cookie, error) {
	args := m.Called(ctx, cookieID, title, category)
	return args.Get(0).(*Cookie), args.Error(1)
}
