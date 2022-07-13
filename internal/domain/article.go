package domain

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Author       uuid.UUID `json:"author"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
	UpdateDate   time.Time `json:"update_date"`
}
