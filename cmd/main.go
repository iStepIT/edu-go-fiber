package main

import (
	"edu-go-fiber/config"
	"edu-go-fiber/internal/home"
	"edu-go-fiber/internal/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()

	app := fiber.New()
	app.Use(recover.New())

	pages.NewPagesHandler(app)
	home.NewHandler(app)

	app.Listen(":3000")
}
