package main

import (
	"github.com/Narutchai01/pathFinder-BE_Go-edition/configport"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":" + configport.Portconfig())
}
