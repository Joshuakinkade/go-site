package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joshuakinkade/go-site/models"
)

type IPhotosRepository interface {
	ListPhotos(offset, limit int) ([]models.Photo, error)
	GetPhoto(id uuid.UUID) (models.Photo, error)
	CreatePhoto(photo models.Photo) (models.Photo, error)
}

type PhotosRepository struct {
	db *pgx.Conn
}

func NewPhotosRepository(db *pgx.Conn) PhotosRepository {
	return PhotosRepository{
		db: db,
	}
}

func (r PhotosRepository) ListPhotos(offset, limit int) ([]models.Photo, error) {
	query := "SELECT id, alt_text, caption, uploaded_at FROM photos WHERE deleted_at IS NULL ORDER BY uploaded_at DESC LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(context.TODO(), query, limit, offset)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	photos := []models.Photo{}
	for rows.Next() {
		photo := models.Photo{}
		err := rows.Scan(&photo.ID, &photo.AltText, &photo.Caption, &photo.UploadedAt)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return nil, nil
}

func GetPhoto(id uuid.UUID) (models.Photo, error) {
	return models.Photo{}, nil
}

func CreatePhoto(photo models.Photo) (models.Photo, error) {
	return models.Photo{}, nil
}
