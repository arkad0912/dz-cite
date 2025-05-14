package main

import (
	"dz/config"
	"dz/internal/pages"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//инициализация конфига
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)

	//Создание Fiber
	app := fiber.New()

	//Инициализация

	app.Use(recover.New()) //мидлвейр, обрабатывает ошибки

	pages.NewPagesHandler(app)

	app.Listen(":3000")
}
