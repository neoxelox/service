package repository

import (
	"context"

	"github.com/rs/xid"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (r *MockRepository) CreateOrUpdate(ctx context.Context, cookie Model) (*Model, error) {
	args := r.Called(ctx, cookie)
	return args.Get(0).(*Model), args.Error(1)
}

func (r *MockRepository) GetByID(ctx context.Context, cookieID xid.ID) (*Model, error) {
	args := r.Called(ctx, cookieID)
	return args.Get(0).(*Model), args.Error(1)
}

func (r *MockRepository) ListAll(ctx context.Context) ([]*Model, error) {
	args := r.Called(ctx)
	return args.Get(0).([]*Model), args.Error(1)
}

func (r *MockRepository) Delete(ctx context.Context, cookieID xid.ID) error {
	args := r.Called(ctx, cookieID)
	return args.Error(0)
}
