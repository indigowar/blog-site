package auth

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNewTokenManager(t *testing.T) {
	validSecret := []byte("hello, world")

	_, err := NewTokenManager(nil, time.Hour, time.Minute)
	if err == nil {
		t.Errorf("NewTokenManager got invalid set of arguments but didn't return the error")
	}

	_, err = NewTokenManager(validSecret, time.Millisecond, time.Hour*24)
	if err == nil {
		t.Errorf("NewTokenManager got invalid set of arguments but didn't return the error")
	}

	_, err = NewTokenManager(validSecret, time.Millisecond, time.Hour*24)
	if err == nil {
		t.Errorf("NewTokenManager got invalid set of arguments but didn't return the error")
	}

	_, err = NewTokenManager(validSecret, time.Hour, time.Second*24)
	if err == nil {
		t.Errorf("NewTokenManager got invalid set of arguments but didn't return the error")
	}
}

func TestTokenManager(t *testing.T) {
	manager, _ := NewTokenManager([]byte("hello, world"), time.Hour, 5*time.Minute)

	id := uuid.New()
	token, err := manager.NewJWT(id)
	if err != nil {
		t.Errorf("TokenManager.NewJWT got an error with valid args:%s", err)
	}

	tokenId, _, err := manager.Verify(token)
	if err != nil {
		t.Errorf("TokenManager.Verify got an error with valid args: %s", err)
	}

	if tokenId != id {
		t.Errorf("The extracted id from token is not the same with input")
	}
}
