package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func LoginController(c *fiber.Ctx) error {

	req := new(Login)

	c.BodyParser(req)

	passwordBytes := []byte(req.PassWord)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return c.Send([]byte(err.Error()))
	}

	return c.JSON(fiber.Map{
		"email":    req.Email,
		"password": hash,
	})

}
