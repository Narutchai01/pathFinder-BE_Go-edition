package routes

import (
	controller "github.com/Narutchai01/pathFinder-BE_Go-edition/controller/User"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(a *fiber.App) {
	route := a.Group("/api/user")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this is User route",
		})
	})

	route.Post("/login", controller.LoginController)
}
