package services

import (
	"github.com/google/uuid"
	"github.com/joshuakinkade/go-site/db"
	"github.com/joshuakinkade/go-site/lib/photos"
	"github.com/joshuakinkade/go-site/models"
)

// PhotoService provides methods for working with photos.
type PhotoService struct {
	photos  db.IPhotosRepository
	library photos.ILibrary
}

// New returns an initialized PhotoService
func New(photos db.IPhotosRepository) PhotoService {
	return PhotoService{
		photos: photos,
	}
}

// UploadPhoto saves a photo to storage and adds it to the db.
func (p PhotoService) UploadPhoto(photo []byte) (uuid.UUID, error) {
	// Are there any adjustments I want to do? Mabye make sure it's oriented correctly?
	// Save to storage
	// Save to repository
	return uuid.New(), nil
}

// ListPhotos returns a list of photos in reverse chronological order
func (p PhotoService) ListPhotos(offset, limit int) ([]models.Photo, error) {
	photos, err := p.photos.ListPhotos(offset, limit)
	return photos, err
}

// GetPhoto returns a photo's content with some optional optimizations.
func (p PhotoService) GetPhoto(id uuid.UUID) ([]byte, error) {
	// Build optimized photo name
	name := id.String() + ""
	data, err := p.library.LoadPhoto(name)
	if err != nil {
		return nil, err
	}
	if data == nil { // optimized photo not found
		data, err = p.library.LoadPhoto(id.String())
		if err != nil {
			return nil, err
		}
		// Resize photo - resarch go image libraries and either use it directly, or wrap it.
		// Scale
		// Crop
		// Adjust quality - at least for jpeg
		// Get Output - jpeg, png, webp, heic
	}
	return data, nil
}
