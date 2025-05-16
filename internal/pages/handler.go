// internal/pages/handler.go
package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) Home(c *fiber.Ctx) error {
	h.logger.Info("Handling home request")
	return c.SendString("Hello, World!")
}

func SetupRoutes(app *fiber.App, logger *slog.Logger) {
	h := NewHandler(logger)
	app.Get("/", h.Home)
}
