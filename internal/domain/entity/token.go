package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Token     string    `json:"token" gorm:"not null;uniqueIndex"`
	Revoked   bool      `json:"revoked" gorm:"not null;default:false"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (rt *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	if rt.ID == uuid.Nil {
		rt.ID = uuid.New()
	}
	return nil
}

func (RefreshToken) TableName() string { return "refresh_tokens" }
