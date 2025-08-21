package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(rt *entity.RefreshToken) error
	Find(token string) (*entity.RefreshToken, error)
	Revoke(token string) error
	RevokeAllForUser(userID uuid.UUID) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(rt *entity.RefreshToken) error { return r.db.Create(rt).Error }

func (r *tokenRepository) Find(token string) (*entity.RefreshToken, error) {
	var t entity.RefreshToken
	if err := r.db.Where("token = ?", token).First(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *tokenRepository) Revoke(token string) error {
	return r.db.Model(&entity.RefreshToken{}).Where("token = ?", token).Update("revoked", true).Error
}

func (r *tokenRepository) RevokeAllForUser(userID uuid.UUID) error {
	return r.db.Model(&entity.RefreshToken{}).Where("user_id = ?", userID).Update("revoked", true).Error
}

// Cleanup helper (optional)
func (r *tokenRepository) DeleteExpired(now time.Time) error {
	return r.db.Where("expires_at < ?", now).Delete(&entity.RefreshToken{}).Error
}
