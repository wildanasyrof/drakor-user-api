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

type FavoriteService interface {
	Create(userID uuid.UUID, req *dto.FavoriteRequest) (*entity.Favorite, error)
	Delete(userID uuid.UUID, favoriteID uuid.UUID) (*entity.Favorite, error)
	GetByUser(userID uuid.UUID) (interface{}, error)
}

type favoriteService struct {
	favoriteRepository repository.FavoriteRepository
	cfg                *config.Config
	logger             logger.Logger
}

func NewFavoriteService(favoriteRepository repository.FavoriteRepository, cfg *config.Config, logger logger.Logger) FavoriteService {
	return &favoriteService{
		favoriteRepository: favoriteRepository,
		cfg:                cfg,
		logger:             logger,
	}
}

// Create implements FavoriteService.
func (f *favoriteService) Create(userID uuid.UUID, req *dto.FavoriteRequest) (*entity.Favorite, error) {
	favorite := &entity.Favorite{
		UserID:    userID,
		DramaSlug: req.DramaSlug,
	}

	if err := f.favoriteRepository.Create(favorite); err != nil {
		return nil, err
	}

	return favorite, nil
}

// Delete implements FavoriteService.
func (f *favoriteService) Delete(userID uuid.UUID, favoriteID uuid.UUID) (*entity.Favorite, error) {
	favorite, err := f.favoriteRepository.GetByID(favoriteID)
	if err != nil {
		return nil, errors.New("favorite not found")
	}

	if err := f.favoriteRepository.Delete(userID, favoriteID); err != nil {
		return nil, err
	}

	return favorite, nil
}

// GetByUser implements FavoriteService.
func (f *favoriteService) GetByUser(userID uuid.UUID) (interface{}, error) {
	favorites, err := f.favoriteRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: 10 * time.Second}
	out := []fiber.Map{}
	for _, it := range favorites {
		url := fmt.Sprintf("%s/drama/%s", f.cfg.Server.ScraperBaseURL, it.DramaSlug)
		resp, err := client.Get(url)
		if err != nil {
			f.logger.Error(err, "Failed to fetch drama details for slug: "+it.DramaSlug+" error: "+err.Error())
			continue
		}
		defer resp.Body.Close()
		var detail map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&detail); err != nil {
			f.logger.Error(err, "Failed to decode response body for slug: "+it.DramaSlug)
			continue
		}
		if data, ok := detail["data"].(map[string]interface{}); ok {
			out = append(out, fiber.Map(data))
		} else {
			continue
		}
	}

	f.logger.Info("Retrieved favorites for user: " + userID.String() + " count: " + fmt.Sprint(len(out)))

	return out, nil
}
