package database

import (
	"api/internal/models"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config models.DatabaseConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.User, config.Password, config.Name, config.Port, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("Failed to connect to database", zap.Error(err))
	}
	return db
}
