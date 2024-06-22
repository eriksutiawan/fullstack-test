package usecase

import (
	userusecase "backend/internal/app/user/usecase"
	"backend/internal/pkg"
	"context"
)

type SubmissionService struct {
	userservice userusecase.IUserService
}

func NewSubmissionService(user userusecase.IUserService) ISubmissionService {
	return &SubmissionService{userservice: user}
}

func (s *SubmissionService) Register(ctx context.Context, dto RegisterDto) error {
	if err := dto.Validate(); err != nil {
		return pkg.BadRequestError
	}
	if err := s.userservice.Save(ctx, userusecase.UserCreatorDto{
		Name:           dto.Name,
		IdentityNumber: dto.IdentityNumber,
		Email:          dto.Email,
		DateOfBirth:    dto.DateOfBirth,
	}); err != nil {
		return err
	}

	return nil
}
