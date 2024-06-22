package usecase

import (
	"backend/internal/app/user/repository"
	"backend/internal/pkg"
	"context"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (s *UserService) Save(ctx context.Context, dto UserCreatorDto) error {
	user := *repository.NewUserModel().
		SetName(dto.Name).
		SetEmail(dto.Email).
		SetIdentityNumber(dto.IdentityNumber).
		SetDateOfBirth(dto.DateOfBirth)

	if err := s.repo.Save(ctx, user); err != nil {
		return pkg.InternalServerError
	}

	return nil
}
