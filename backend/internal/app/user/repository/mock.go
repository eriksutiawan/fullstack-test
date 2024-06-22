package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(ctx context.Context, user UserModel) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
