package repository

import (
	"context"
	"rest_websocket/models"
)

// Pattern repository
// Defined Interface
type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	//cerra conexion a base de datos
	Close() error
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
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

// funcion para cerrar la conexion
func Close() error {
	//devuelve lo que la implementacion de la interfaz retorne
	return implementation.Close()
}
