package main

import (
	"github.com/Narutchai01/pathFinder-BE_Go-edition/config"
	"github.com/Narutchai01/pathFinder-BE_Go-edition/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello",
		})
	})

	routes.UserRouter(app)

	app.Listen(":" + config.Portconfig())
}
