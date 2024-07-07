package models

import "github.com/google/uuid"

type Photo struct {
	ID      uuid.UUID `json:"id"`
	AltText string    `json:"altText"`
}

// Resize resizes the photo to the specified height, width, and quality.
func (photo Photo) Resize(height, width, quality int) ([]byte, error) {
	return nil, nil
}
