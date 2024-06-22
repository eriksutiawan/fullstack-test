package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Save(ctx context.Context, dto UserCreatorDto) error {
	args := m.Called(ctx, dto)
	return args.Error(0)
}
