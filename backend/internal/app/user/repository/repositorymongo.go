package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(database *mongo.Database) IUserRepository {
	collection := database.Collection("users")
	return &UserRepository{collection: collection}
}

func (r *UserRepository) Save(ctx context.Context, model UserModel) error {
	_, err := r.collection.InsertOne(ctx, model)
	return err
}
