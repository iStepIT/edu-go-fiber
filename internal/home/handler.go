package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")

	api.Get("/test", h.test)
}

func (h *HomeHandler) test(c *fiber.Ctx) error {
	users := []User{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jack"},
		{Id: 3, Name: "Smith"},
	}
	names := []string{"Tom", "Jack", "Smith"}
	data := struct {
		Names []string
		Users []User
	}{names, users}
	return c.Render("page", data)
}
