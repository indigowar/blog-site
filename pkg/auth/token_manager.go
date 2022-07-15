package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type TokenManager interface {
	NewJWT(uuid.UUID) (string, error)
	NewRefresh(uuid.UUID) string
	Verify(string) (uuid.UUID, bool, error)
}

type tokenManager struct {
	signingKey []byte
	refreshExp time.Duration
	accessExp  time.Duration
}

func (manager *tokenManager) NewJWT(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   id.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.accessExp)),
		Issuer:    "token_manager",
	})

	tokenString, err := token.SignedString(manager.signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (manager *tokenManager) NewRefresh(_ uuid.UUID) string {
	// TODO: Implement creating of the refresh token
	panic("not implemented")
}

func (manager *tokenManager) Verify(tokenString string) (uuid.UUID, bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return manager.signingKey, nil
	})

	if err != nil {
		return uuid.UUID{}, false, errors.New("failed to parse token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, false, errors.New("failed to get claims")
	}

	stringId, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, false, errors.New("failed to get id")
	}

	id, err := uuid.Parse(stringId)
	if err != nil {
		return uuid.UUID{}, false, errors.New("failed to parse id")
	}

	return id, claims.VerifyExpiresAt(time.Now().UnixNano(), true), nil
}

func NewTokenManager(secret []byte, refExp, accExp time.Duration) (TokenManager, error) {
	if secret == nil || refExp < time.Hour || accExp < time.Second*30 || accExp > time.Hour || refExp < accExp {
		return nil, errors.New("the arguments is wrong")
	}
	return &tokenManager{secret, refExp, accExp}, nil
}
