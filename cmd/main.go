package main

import (
	"dz/config"
	"dz/internal/pages"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	//инициализация конфига
	config.Init()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)

	// Создание Fiber приложения
	app := fiber.New()

	// Подключение middleware
	app.Use(recover.New())         // Мидлвейр, обрабатывающий ошибки
	app.Use(slogfiber.New(logger)) // Логирование запросов

	//pages.NewPagesHandler(app)
	app.Get("/", pages.Handler) // Обработчик для корневого пути

	app.Listen(":3000")
}
