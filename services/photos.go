package services

import (
	"github.com/google/uuid"
	"github.com/joshuakinkade/go-site/db"
	"github.com/joshuakinkade/go-site/models"
)

// Depenenencies:
// photos repository
// s3/storage library
// image processing library

type PhotoService struct {
	photos db.IPhotosRepository
}

func NewPhotoService(photos db.IPhotosRepository) PhotoService {
	return PhotoService{
		photos: photos,
	}
}

func (p PhotoService) UploadPhoto(photo []byte) (uuid.UUID, error) {
	// Are there any adjustments I want to do? Mabye make sure it's oriented correctly?
	// Save to storage
	// Save to repository
	return uuid.New(), nil
}

func (p PhotoService) ListPhotos(offset, limit int) ([]models.Photo, error) {
	photos, err := p.photos.ListPhotos(offset, limit)
	return photos, err
}

func (p PhotoService) GetPhoto(id uuid.UUID) ([]byte, error) {
	// Look up photo in repository
	// Look for photo in storage
	// Resize photo
	return nil, nil
}
