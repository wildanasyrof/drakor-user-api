package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Favorite struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index;uniqueIndex:uni_favorites_user_slug"`
	DramaSlug string    `json:"drama_slug" gorm:"not null;uniqueIndex:uni_favorites_user_slug"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Optional relation; creates FK with ON DELETE CASCADE:
	User User `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (f *Favorite) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

func (Favorite) TableName() string { return "favorites" }
