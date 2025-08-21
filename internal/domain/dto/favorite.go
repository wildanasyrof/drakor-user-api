package dto

type FavoriteRequest struct {
	DramaSlug string `json:"drama_slug" validate:"required"`
}

type FavoriteResponse struct {
	DramaSlug string `json:"drama_slug"`
}
