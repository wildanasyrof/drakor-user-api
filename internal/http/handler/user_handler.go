package handler

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/service"
	"github.com/wildanasyrof/drakor-user-api/pkg/response"
	"github.com/wildanasyrof/drakor-user-api/pkg/storage"
	"github.com/wildanasyrof/drakor-user-api/pkg/validator"
)

type UserHandler struct {
	userService service.UserService
	validator   validator.Validator
	storage     storage.LocalStorage
}

func NewUserHandler(userService service.UserService, validator validator.Validator, storage storage.LocalStorage) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
		storage:     storage,
	}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	user, err := h.userService.Get(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get user profile", err)
	}
	return response.Success(c, "User profile retrieved successfully", user)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", err)
	}

	userID := c.Locals("user_id").(uuid.UUID)
	user, err := h.userService.Update(userID, &req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update user profile", err)
	}
	return response.Success(c, "User profile updated successfully", user)
}

func (h *UserHandler) UpdateAvatar(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid file upload", err)
	}

	if file.Size > 10*1024*1024 {
		return response.Error(c, fiber.StatusBadRequest, "file too large", nil)
	}
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		ext = ".bin"
	}

	filename, err := h.storage.Save(file)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to save file", err)
	}

	url := "/uploads/" + filename
	userID := c.Locals("user_id").(uuid.UUID)
	user, err := h.userService.UpdateAvatar(userID, url)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to update avatar", err.Error())
	}
	return response.Success(c, "Avatar updated successfully", user)
}
