package services

import (
	"github.com/google/uuid"
)

// Depenenencies:
// photos repository
// s3/storage library
// image processing library

type PhotoService struct {}

func NewPhotoService() PhotoService {
	return PhotoService{}
}

func (p PhotoService) GetPhoto(id uuid.UUID) ([]byte, error) {
	return nil, nil
}
