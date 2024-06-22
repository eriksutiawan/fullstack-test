package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	id             primitive.ObjectID `bson:"_id,omitempty"`
	name           string             `bson:"name"`
	identityNumber string             `bson:"identity_number"`
	email          string             `bson:"email"`
	dateOfBirth    string             `bson:"date_of_birth"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

// Getters
func (r *UserModel) GetId() primitive.ObjectID {
	return r.id
}

func (r *UserModel) SetId(id primitive.ObjectID) *UserModel {
	r.id = id
	return r
}

func (r *UserModel) GetName() string {
	return r.name
}

func (r *UserModel) GetIdentityNumber() string {
	return r.identityNumber
}

func (r *UserModel) GetEmail() string {
	return r.email
}

func (r *UserModel) GetDateOfBirth() string {
	return r.dateOfBirth
}

// Setters
func (r *UserModel) SetName(name string) *UserModel {
	r.name = name
	return r
}

func (r *UserModel) SetIdentityNumber(identityNumber string) *UserModel {
	r.identityNumber = identityNumber
	return r
}

func (r *UserModel) SetEmail(email string) *UserModel {
	r.email = email
	return r
}

func (r *UserModel) SetDateOfBirth(dateOfBirth string) *UserModel {
	r.dateOfBirth = dateOfBirth
	return r
}
