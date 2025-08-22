package dto

import (
	"time"

	"github.com/google/uuid"
)

type FavoriteRequest struct {
	DramaSlug string `json:"drama_slug" validate:"required"`
}

type FavoriteResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	DramaSlug string    `json:"drama_slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
