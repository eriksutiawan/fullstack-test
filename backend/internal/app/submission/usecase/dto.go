package usecase

import (
	"github.com/go-playground/validator"
)

type RegisterDto struct {
	Name           string `json:"name" validate:"required"`
	IdentityNumber string `json:"identity_number" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	DateOfBirth    string `json:"date_of_birth" validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (v RegisterDto) Validate() error {
	if err := validate.Struct(v); err != nil {
		return err
	}
	return nil
}
