package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/service"
	"github.com/wildanasyrof/drakor-user-api/pkg/response"
	"github.com/wildanasyrof/drakor-user-api/pkg/validator"
)

type HistoryHandler struct {
	historyService service.HistoryService
	validator      validator.Validator
}

func NewHistoryHandler(historyService service.HistoryService, validator validator.Validator) *HistoryHandler {
	return &HistoryHandler{
		historyService: historyService,
		validator:      validator,
	}
}

func (h *HistoryHandler) Create(c *fiber.Ctx) error {
	var req dto.HistoryRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	userID := c.Locals("user_id").(uuid.UUID)
	history, err := h.historyService.Create(userID, &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create history", err)
	}
	return response.Success(c, "History created successfully", history)
}

func (h *HistoryHandler) Delete(c *fiber.Ctx) error {
	historyID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid history ID", err)
	}

	userID := c.Locals("user_id").(uuid.UUID)
	history, err := h.historyService.Delete(userID, historyID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete history", err.Error())
	}

	return response.Success(c, "History deleted successfully", history)
}

func (h *HistoryHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	histories, err := h.historyService.GetByUser(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve histories", err.Error())
	}

	return response.Success(c, "Histories retrieved successfully", histories)
}

func (h *HistoryHandler) Update(c *fiber.Ctx) error {
	var req dto.UpdateHistoryRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	historyID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid history ID", err)
	}

	userID := c.Locals("user_id").(uuid.UUID)
	history, err := h.historyService.Update(userID, historyID, &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update history", err)
	}

	return response.Success(c, "History updated successfully", history)
}
