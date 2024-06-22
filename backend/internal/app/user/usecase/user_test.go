package usecase_test

import (
	"backend/internal/app/user/repository"
	"backend/internal/app/user/usecase"
	"backend/internal/pkg"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Save(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := usecase.NewUserService(mockRepo)

	t.Run("should save user successfully", func(t *testing.T) {
		dto := usecase.UserCreatorDto{
			Name:           "John Doe",
			Email:          "john.doe@example.com",
			IdentityNumber: "123456789",
			DateOfBirth:    "2000-01-01",
		}

		mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil).Once()

		err := service.Save(context.Background(), dto)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return internal server error on repository failure", func(t *testing.T) {
		dto := usecase.UserCreatorDto{
			Name:           "John Doe",
			Email:          "john.doe@example.com",
			IdentityNumber: "123456789",
			DateOfBirth:    "2000-01-01",
		}

		mockRepo.On("Save", mock.Anything, mock.Anything).Return(errors.New("repository error")).Once()

		err := service.Save(context.Background(), dto)

		assert.Equal(t, pkg.InternalServerError, err)
		mockRepo.AssertExpectations(t)
	})
}
