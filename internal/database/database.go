package database

import (
	"github.com/tuts-4-all/backend/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
