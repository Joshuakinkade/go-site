package models

import "github.com/google/uuid"

type Photo struct {
	ID      uuid.UUID `json:"id"`
	AltText string    `json:"altText"`
	Caption string    `json:"caption"`
}

// Resize resizes the photo to the specified height, width, and quality.
// Does this to file i/o? Do I want that here or in a library?
// 1. Load the image from storage
// 2. Resize the image
// 3. Save the image to storage
func (photo Photo) Resize(height, width, quality int) ([]byte, error) {
	return nil, nil
}
