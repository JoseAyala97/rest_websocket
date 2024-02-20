package repository

import (
	"context"
	"rest_websocket/models"
)

// Pattern repository
// Defined Interface
type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
}

var implementation UserRepository

// Method
func SetRepository(repository UserRepository) {
	implementation = repository
}

// Method insert new user
func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

// Method GetUserById
func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}
