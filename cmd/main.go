package main

import (
	"edu-go-fiber/config"
	"edu-go-fiber/internal/home"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Printf("dbConf: %v", dbConf)
	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")
}
