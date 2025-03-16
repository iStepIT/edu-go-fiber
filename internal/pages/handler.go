package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router     fiber.Router
	slogLogger *slog.Logger
}

func NewPagesHandler(router fiber.Router, slogLogger *slog.Logger) {
	h := &PagesHandler{
		router:     router,
		slogLogger: slogLogger,
	}
	api := h.router.Group("/api")

	api.Get("/home", h.homePage)
}

func (h *PagesHandler) homePage(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
