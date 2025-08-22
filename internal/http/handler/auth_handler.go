package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/service"
	"github.com/wildanasyrof/drakor-user-api/pkg/response"
	"github.com/wildanasyrof/drakor-user-api/pkg/validator"
)

type AuthHandler struct {
	authService service.AuthService
	validator   validator.Validator
}

func NewAuthHandler(authService service.AuthService, validator validator.Validator) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation error", err)
	}

	user, token, err := h.authService.Register(&req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Registration failed", err)
	}

	return response.Success(c, "User registered successfully", fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation error", err)
	}

	user, token, err := h.authService.Login(&req)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Login failed", err.Error())
	}

	return response.Success(c, "User logged in successfully", fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req dto.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation error", err)
	}

	newToken, err := h.authService.Refresh(req.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Refresh token failed", err)
	}

	return response.Success(c, "Token refreshed successfully", fiber.Map{
		"access_token": newToken,
	})
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var req dto.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation error", err)
	}

	newToken, err := h.authService.Refresh(req.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Refresh token failed", err.Error())
	}

	return response.Success(c, "Token refreshed successfully", fiber.Map{
		"access_token": newToken,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var req dto.LogoutUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.ValidateBody(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Validation error", err)
	}

	err := h.authService.Logout(req.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Logout failed", err)
	}

	return response.Success(c, "User logged out successfully", nil)
}
