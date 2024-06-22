package usecase

type UserCreatorDto struct {
	Name           string `json:"name" validate:"required"`
	IdentityNumber string `json:"identity_number" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	DateOfBirth    string `json:"date_of_birth" validate:"required"`
}
