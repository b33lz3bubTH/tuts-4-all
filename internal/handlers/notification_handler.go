package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tuts-4-all/backend/internal/database"
	"github.com/tuts-4-all/backend/internal/models"
	"strconv"
)

type NotificationHandler struct {
	repo database.NotificationRepository
}

func NewNotificationHandler(repo database.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repo: repo}
}

func (h *NotificationHandler) RegisterRoutes(app *fiber.App) {
	notifications := app.Group("/notifications")
	notifications.Get("/", h.List)
	notifications.Get("/:id", h.Get)
	notifications.Post("/", h.Create)
	notifications.Put("/:id", h.Update)
	notifications.Delete("/:id", h.Delete)
}

func (h *NotificationHandler) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	
	offset := (page - 1) * limit
	items, total, err := h.repo.List(offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	return c.JSON(fiber.Map{
		"data": items,
		"total": total,
		"page": page,
		"limit": limit,
		"pages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *NotificationHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	n, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Notification not found",
		})
	}

	return c.JSON(n)
}

func (h *NotificationHandler) Create(c *fiber.Ctx) error {
	var n models.Notification
	if err := c.BodyParser(&n); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.repo.Create(&n); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(n)
}

func (h *NotificationHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	existing, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Notification not found",
		})
	}

	var updates models.Notification
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	existing.Title = updates.Title
	existing.Content = updates.Content

	if err := h.repo.Update(existing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(existing)
}

func (h *NotificationHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
