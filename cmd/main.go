package main

import (
	"edu-go-fiber/config"
	"edu-go-fiber/internal/home"
	"edu-go-fiber/internal/pages"
	"edu-go-fiber/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
	"log/slog"
	"os"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConf := config.NewLogConfig()
	customLogger := logger.NewLogger(logConf)
	slogLogger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	app := fiber.New()
	app.Use(slogfiber.New(slogLogger))
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	pages.NewPagesHandler(app, slogLogger)
	home.NewHandler(app, customLogger)

	app.Listen(":3000")
}
