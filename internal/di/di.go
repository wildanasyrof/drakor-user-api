package di

import (
	"github.com/wildanasyrof/drakor-user-api/internal/config"
	"github.com/wildanasyrof/drakor-user-api/internal/db"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
	"gorm.io/gorm"
)

type DI struct {
	logger logger.Logger
	db     *gorm.DB
}

func InitDI(cfg *config.Config) *DI {

	logger := logger.NewZerologLogger(cfg.Server.Env)
	db := db.Connect(cfg, logger)

	return &DI{
		logger: logger,
		db:     db,
	}
}
