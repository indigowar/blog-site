package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/indigowar/blog-site/internal/domain"
)

var (
	ErrUserAlreadyExists = errors.New("error: user already exists")
	ErrUserNotFound      = errors.New("error: user was not found")
)

type UserRepository interface {
	Create(user domain.User) error
	Delete(uuid uuid.UUID) error
	Get(uuid.UUID) (domain.User, error)
	ChangeEmail(uuid.UUID, string) error
	ChangePassword(uuid.UUID, string) error
}
