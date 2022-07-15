package auth

import (
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
