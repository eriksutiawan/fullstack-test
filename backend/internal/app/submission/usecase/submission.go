package usecase

import (
	"context"
)

type ISubmissionService interface {
	Register(ctx context.Context, dto RegisterDto) error
}
