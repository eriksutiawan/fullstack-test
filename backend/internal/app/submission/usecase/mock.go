package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// Mock implementation of ISubmissionService
type MockSubmissionService struct {
	mock.Mock
}

func (m *MockSubmissionService) Register(ctx context.Context, dto RegisterDto) error {
	args := m.Called(ctx, dto)
	return args.Error(0)
}
