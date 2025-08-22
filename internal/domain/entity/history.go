package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type History struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index;uniqueIndex:uni_histories_user_slug"`
	DramaSlug      string    `json:"drama_slug" gorm:"not null;uniqueIndex:uni_histories_user_slug"`
	DramaEps       int64     `json:"drama_eps" gorm:"not null"` // use int64 if you want BIGINT
	WatchedSeconds *int64    `json:"watched_seconds"`           // nullable
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User User `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (h *History) BeforeCreate(tx *gorm.DB) (err error) {
	if h.ID == uuid.Nil {
		h.ID = uuid.New()
	}
	return nil
}

func (History) TableName() string { return "histories" }
