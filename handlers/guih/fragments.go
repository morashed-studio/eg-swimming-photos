package guih

import (
	"context"
	"goweb/ui/fragments"

	"github.com/gofiber/fiber/v2"
)

func DashboardFragment(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  fragments.Dashboard().Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(200)
}
