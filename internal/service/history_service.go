package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/config"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"github.com/wildanasyrof/drakor-user-api/internal/repository"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
)

type HistoryService interface {
	Create(userID uuid.UUID, req *dto.HistoryRequest) (*entity.History, error)
	Update(userID uuid.UUID, historyID uuid.UUID, req *dto.UpdateHistoryRequest) (*entity.History, error)
	Delete(userID uuid.UUID, historyID uuid.UUID) (*entity.History, error)
	GetByUser(userID uuid.UUID) ([]dto.HistoryResponse, error)
}

type historyService struct {
	historyRepo repository.HistoryRepository
	cfg         *config.Config
	logger      logger.Logger
}

func NewHistoryService(historyRepo repository.HistoryRepository, cfg *config.Config, logger logger.Logger) HistoryService {
	return &historyService{
		historyRepo: historyRepo,
		cfg:         cfg,
		logger:      logger,
	}
}

// Create implements HistoryService.
func (h *historyService) Create(userID uuid.UUID, req *dto.HistoryRequest) (*entity.History, error) {
	history := &entity.History{
		UserID:         userID,
		DramaSlug:      req.DramaSlug,
		DramaEps:       req.Eps,
		WatchedSeconds: req.WatchedSeconds,
	}

	if err := h.historyRepo.Create(history); err != nil {
		return nil, err
	}

	return history, nil
}

// Delete implements HistoryService.
func (h *historyService) Delete(userID uuid.UUID, historyID uuid.UUID) (*entity.History, error) {
	history, err := h.historyRepo.GetByID(userID, historyID)
	if err != nil {
		return nil, errors.New("history not found")
	}

	if err := h.historyRepo.DeleteByID(history.ID); err != nil {
		return nil, err
	}

	return history, nil
}

// GetByUserID implements HistoryService.
func (h *historyService) GetByUser(userID uuid.UUID) ([]dto.HistoryResponse, error) {
	histories, err := h.historyRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: 10 * time.Second}
	out := make([]dto.HistoryResponse, 0, len(histories))

	for _, it := range histories {
		// IMPORTANT: copy the range item before taking its address
		hist := it

		url := fmt.Sprintf("%s/drama/%s", h.cfg.Server.ScraperBaseURL, hist.DramaSlug)
		resp, err := client.Get(url)
		if err != nil {
			h.logger.Error(err, "Failed to fetch drama details for slug: "+hist.DramaSlug+" error: "+err.Error())
			// Include history with empty detail to keep shape stable
			out = append(out, dto.HistoryResponse{History: &hist, Detail: map[string]interface{}{}})
			continue
		}

		// Do NOT defer in a loop; close immediately after use
		var payload map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
			resp.Body.Close()
			h.logger.Error(err, "Failed to decode response body for slug: "+hist.DramaSlug)
			out = append(out, dto.HistoryResponse{History: &hist, Detail: map[string]interface{}{}})
			continue
		}
		resp.Body.Close()

		// Extract the "data" field from the scraper response
		detail, _ := payload["data"].(map[string]interface{})
		if detail == nil {
			detail = map[string]interface{}{}
		}

		out = append(out, dto.HistoryResponse{
			History: &hist,
			Detail:  fiber.Map(detail), // optional: normalize to fiber.Map
		})
	}

	h.logger.Info("Retrieved histories for user: " + userID.String() + " count: " + fmt.Sprint(len(out)))
	return out, nil
}

// Update implements HistoryService.
func (h *historyService) Update(userID uuid.UUID, historyID uuid.UUID, req *dto.UpdateHistoryRequest) (*entity.History, error) {
	history, err := h.historyRepo.GetByID(userID, historyID)
	if err != nil {
		return nil, err
	}

	history.ID = historyID
	history.DramaEps = req.Eps
	history.WatchedSeconds = req.WatchedSeconds

	if err := h.historyRepo.Update(history); err != nil {
		return nil, err
	}
	return history, nil
}
