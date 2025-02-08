package guih

import (
	"context"

	anc "goweb/ancillaries"
	"goweb/db/sections"
	"goweb/ui/fragments"

	"github.com/gofiber/fiber/v2"
)

func DashboardFragment(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  allSections := anc.Must(sections.GetAll()).([]sections.DataModel)
  fragments.Dashboard(allSections).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(fiber.StatusOK)
}
