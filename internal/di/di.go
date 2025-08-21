package di

import (
	"github.com/wildanasyrof/drakor-user-api/internal/config"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
)

type DI struct {
	logger logger.Logger
}

func InitDI(cfg *config.Config) *DI {

	logger := logger.NewZerologLogger(cfg.Server.Env)

	return &DI{
		logger: logger,
	}
}
