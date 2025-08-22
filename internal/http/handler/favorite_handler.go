package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/service"
	"github.com/wildanasyrof/drakor-user-api/pkg/response"
	"github.com/wildanasyrof/drakor-user-api/pkg/validator"
)

type FavoriteHandler struct {
	favoriteService service.FavoriteService
	validator       validator.Validator
}

func NewFavoriteHandler(favoriteService service.FavoriteService, validator validator.Validator) *FavoriteHandler {
	return &FavoriteHandler{
		favoriteService: favoriteService,
		validator:       validator,
	}
}

func (h *FavoriteHandler) Create(c *fiber.Ctx) error {
	var req dto.FavoriteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	userID := c.Locals("user_id").(uuid.UUID)
	favorite, err := h.favoriteService.Create(userID, &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create favorite", err)
	}

	return response.Success(c, "Favorite created successfully", favorite)
}

func (h *FavoriteHandler) Delete(c *fiber.Ctx) error {
	favoriteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid favorite ID", err)
	}

	userID := c.Locals("user_id").(uuid.UUID)
	favorite, err := h.favoriteService.Delete(userID, favoriteID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete favorite", err.Error())
	}

	return response.Success(c, "Favorite deleted successfully", favorite)
}

func (h *FavoriteHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	favorites, err := h.favoriteService.GetByUser(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve favorites", err.Error())
	}

	return response.Success(c, "Favorites retrieved successfully", favorites)
}
