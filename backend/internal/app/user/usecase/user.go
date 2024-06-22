package usecase

import (
	"context"
)

type IUserService interface {
	Save(ctx context.Context, dto UserCreatorDto) error
}
