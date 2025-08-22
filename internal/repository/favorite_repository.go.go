package repository

import (
	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"gorm.io/gorm"
)

type FavoriteRepository interface {
	Create(favorite *entity.Favorite) error
	GetByID(favoriteID uuid.UUID) (*entity.Favorite, error)
	Delete(userID uuid.UUID, favoriteID uuid.UUID) error
	GetByUserID(userID uuid.UUID) ([]entity.Favorite, error)
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{db: db}
}

// Create implements FavoriteRepository.
func (f *favoriteRepository) Create(favorite *entity.Favorite) error {
	return f.db.Create(favorite).Error
}

// Delete implements FavoriteRepository.
func (f *favoriteRepository) Delete(userID uuid.UUID, favoriteID uuid.UUID) error {
	return f.db.Where("user_id = ? AND id = ?", userID, favoriteID).Delete(&entity.Favorite{}).Error
}

// GetByUserID implements FavoriteRepository.
func (f *favoriteRepository) GetByUserID(userID uuid.UUID) ([]entity.Favorite, error) {
	var favorites []entity.Favorite
	err := f.db.Where("user_id = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

// GetByID implements FavoriteRepository.
func (f *favoriteRepository) GetByID(favoriteID uuid.UUID) (*entity.Favorite, error) {
	var favorite entity.Favorite
	err := f.db.Where("id = ?", favoriteID).First(&favorite).Error
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}
