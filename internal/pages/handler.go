package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}
	api := h.router.Group("/api")

	api.Get("/home", h.homePage)
}

func (h *PagesHandler) homePage(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
