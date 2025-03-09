package main

import (
	"edu-go-fiber/config"
	"edu-go-fiber/internal/home"
	"edu-go-fiber/internal/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConf := config.NewLogConfig()

	app := fiber.New()
	log.SetLevel(log.Level(logConf.Level))

	app.Use(logger.New())
	app.Use(recover.New())

	pages.NewPagesHandler(app)
	home.NewHandler(app)

	app.Listen(":3000")
}
