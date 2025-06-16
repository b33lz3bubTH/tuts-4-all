package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func requestIDMiddleware(c *fiber.Ctx) error {
	rid := c.Get("X-Request-ID")
	if rid == "" {
		rid = uuid.New().String()
		c.Set("X-Request-ID", rid)
	}
	c.Locals("requestid", rid)
	return c.Next()
}

func loggingMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rid := c.Locals("requestid")
		logger.Info("request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Any("request_id", rid),
		)
		return c.Next()
	}
}
