package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")

	api.Get("/error", h.error)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {

	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("method", c.Method()).
		Int("status", fiber.StatusOK).
		Msg("Info")

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("Info")
	return c.SendString("Error")
}
