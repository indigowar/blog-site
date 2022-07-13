package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	RegistrationDate time.Time `json:"registration_date"`
}
