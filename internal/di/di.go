package di

import (
	"github.com/wildanasyrof/drakor-user-api/internal/config"
	"github.com/wildanasyrof/drakor-user-api/internal/db"
	"github.com/wildanasyrof/drakor-user-api/internal/http/handler"
	"github.com/wildanasyrof/drakor-user-api/internal/repository"
	"github.com/wildanasyrof/drakor-user-api/internal/service"
	"github.com/wildanasyrof/drakor-user-api/pkg/jwt"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
	"github.com/wildanasyrof/drakor-user-api/pkg/validator"
	"gorm.io/gorm"
)

type DI struct {
	logger      logger.Logger
	db          *gorm.DB
	JWT         jwt.JWTService
	AuthHandler *handler.AuthHandler
}

func InitDI(cfg *config.Config) *DI {

	logger := logger.NewZerologLogger(cfg.Server.Env)
	db := db.Connect(cfg, logger)
	jwt := jwt.NewJWTService(cfg)
	validator := validator.NewValidator()

	tokenRepo := repository.NewTokenRepository(db)

	authRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(authRepo, tokenRepo, jwt)
	authHandler := handler.NewAuthHandler(authService, validator)

	return &DI{
		logger:      logger,
		db:          db,
		JWT:         jwt,
		AuthHandler: authHandler,
	}
}
