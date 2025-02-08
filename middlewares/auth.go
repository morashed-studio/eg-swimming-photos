package middlewares

import (
	"eg-swimming-photos/db/users"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	if c.Cookies("username") == "" {
		return c.Redirect("/login")
	}
	data, err := users.Get(c.Cookies("username"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if c.Cookies("password") != data.Password {
		return c.Redirect("/login")
	}
	return c.Next()
}
