package main

import (
	"github.com/tuts-4-all/backend/internal/config"
	"github.com/tuts-4-all/backend/internal/database"
	"github.com/tuts-4-all/backend/internal/handlers"
	"github.com/tuts-4-all/backend/internal/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			logger.NewLogger,
			database.NewDatabase,
			database.NewNotificationRepository,
			handlers.NewFiberApp,
		),
		fx.Invoke(database.Migrate),
		fx.Invoke(handlers.RegisterRoutes),
	).Run()
}
