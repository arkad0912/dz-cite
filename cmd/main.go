package main

import (
	"dz/config"
	"dz/internal/pages"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
}

func setupApp(logger *slog.Logger) *fiber.App {
	app := fiber.New()

	// Middleware в правильном порядке
	app.Use(recover.New())
	app.Use(slogfiber.New(logger))

	// Инициализация роутов
	pages.SetupRoutes(app, logger)

	return app
}

func main() {
	// Инициализация конфигурации
	config.Init()

	// Настройка логгера
	logger := setupLogger()

	// Создание приложения
	app := setupApp(logger)

	app.Listen(":3000")
}
