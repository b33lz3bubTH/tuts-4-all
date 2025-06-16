package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/tuts-4-all/backend/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20240616_create_notifications",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Notification{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("notifications")
			},
		},
	})
	return m.Migrate()
}
