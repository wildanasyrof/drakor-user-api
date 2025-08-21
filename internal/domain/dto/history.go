package dto

type HistoryRequest struct {
	DramaSlug      string `json:"drama_slug" validate:"required"`
	Eps            int64  `json:"eps" validate:"required,min=1"`
	WatchedSeconds *int64 `json:"time_watched"`
}

type HistoryResponse struct {
	DramaSlug      string `json:"drama_slug"`
	Eps            int64  `json:"eps"`
	WatchedSeconds *int64 `json:"time_watched"`
}
