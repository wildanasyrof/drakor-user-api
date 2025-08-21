package entity

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `json:"user_id" gorm:"not null"`
	DramaSlug      string    `json:"drama_slug" gorm:"not null;unique"`
	DramaEps       int       `json:"drama_eps" gorm:"not null"`
	WatchedSeconds int       `json:"watched_seconds"` //`// This field tracks the user's progress
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
