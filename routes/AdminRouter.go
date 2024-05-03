package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AdminRouter(a *fiber.App) {
	route := a.Group("/api/admin")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this is admin route",
		})
	})
}
