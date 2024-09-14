package lib

import "github.com/google/uuid"

// Depenenencies:
// s3 client probably

type PhotoLibrary struct{}

func NewPhotoLibrary() PhotoLibrary {
	return PhotoLibrary{}
}

func (p PhotoLibrary) LoadPhoto(id uuid.UUID) ([]byte, error) {
	return nil, nil
}
