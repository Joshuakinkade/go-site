package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID         uuid.UUID `json:"id"`
	AltText    string    `json:"altText"`
	Caption    string    `json:"caption"`
	UploadedAt time.Time `json:"uploaded_at"`
}
