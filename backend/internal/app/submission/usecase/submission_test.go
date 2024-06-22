// usecase_test.go

package usecase_test

import (
	"backend/internal/app/submission/usecase"
	userusecase "backend/internal/app/user/usecase"
	"backend/internal/pkg"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmissionService_Register(t *testing.T) {
	// Setup mock user service
	mockUserService := new(userusecase.MockUserService)
	submissionService := usecase.NewSubmissionService(mockUserService)

	// Test case 1: Valid registration
	t.Run("Valid registration", func(t *testing.T) {
		ctx := context.Background()
		dto := usecase.RegisterDto{
			Name:           "John Doe",
			IdentityNumber: "123456789",
			Email:          "john.doe@example.com",
			DateOfBirth:    "2000-01-01",
		}

		mockUserService.On("Save", ctx, mock.Anything).Return(nil)

		err := submissionService.Register(ctx, dto)

		assert.NoError(t, err)
		mockUserService.AssertExpectations(t)
	})

	// Test case 2: Invalid registration (DTO validation fails)
	t.Run("Invalid registration (DTO validation fails)", func(t *testing.T) {
		ctx := context.Background()
		dto := usecase.RegisterDto{} // Empty DTO which will fail validation

		err := submissionService.Register(ctx, dto)

		assert.Equal(t, pkg.BadRequestError, err)
	})

}
