package repository

import (
	"context"
)

type IUserRepository interface {
	Save(ctx context.Context, model UserModel) error
}
