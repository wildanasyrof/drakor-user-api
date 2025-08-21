package db

import (
	"github.com/wildanasyrof/drakor-user-api/internal/config"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"github.com/wildanasyrof/drakor-user-api/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config, logger logger.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Database.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect database error, " + err.Error())
	}
	if err := db.AutoMigrate(&entity.User{}, &entity.RefreshToken{}, &entity.Favorite{}, &entity.History{}); err != nil {
		logger.Fatal("auto-migrate failed: " + err.Error())
	}
	return db
}
