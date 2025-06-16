package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tuts-4-all/backend/internal/config"
	"github.com/tuts-4-all/backend/internal/database"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type FiberAppParams struct {
	App                 *fiber.App
	NotificationHandler *NotificationHandler
}

func NewFiberApp(logger *zap.Logger, repo database.NotificationRepository) *fiber.App {
	app := fiber.New()
	app.Use(requestIDMiddleware)
	app.Use(loggingMiddleware(logger))
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	// Register notification routes
	NewNotificationHandler(repo).RegisterRoutes(app)
	return app
}

func RegisterRoutes(lc fx.Lifecycle, app *fiber.App, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go app.Listen(fmt.Sprintf(":%s", cfg.ServerPort))
			return nil
		},
		OnStop: func(context.Context) error {
			return app.Shutdown()
		},
	})
}
