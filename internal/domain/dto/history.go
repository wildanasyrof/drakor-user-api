package dto

import "github.com/wildanasyrof/drakor-user-api/internal/domain/entity"

type HistoryRequest struct {
	DramaSlug      string `json:"slug" validate:"required"`
	Eps            int64  `json:"eps" validate:"required,min=1"`
	WatchedSeconds *int64 `json:"time_watched"`
}

type UpdateHistoryRequest struct {
	Eps            int64  `json:"eps" validate:"required,min=1"`
	WatchedSeconds *int64 `json:"time_watched"`
}

type HistoryResponse struct {
	History *entity.History        `json:"history"`
	Detail  map[string]interface{} `json:"detail"`
}
