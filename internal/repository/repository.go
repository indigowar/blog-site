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

var (
	ErrArticleAlreadyExists = errors.New("error: article already exists")
	ErrArticleNotFound      = errors.New("error: article was not found")
)

type ArticleRepository interface {
	Get(uuid.UUID) (domain.Article, error)
	Create(domain.Article) error
	Delete(uuid.UUID) error
	Update(uuid.UUID, string, string) error
}
