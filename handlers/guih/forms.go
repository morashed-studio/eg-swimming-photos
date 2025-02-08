package guih

import (
	"context"
	anc "goweb/ancillaries"
	"goweb/db/sections"
	"goweb/ui/forms"

	"github.com/gofiber/fiber/v2"
)

func AddSectionForm(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  allSections := anc.Must(sections.GetAll()).([]sections.DataModel)
  forms.AddSection(allSections, nil).Render(
    context.Background(),
    c.Response().BodyWriter(),
  )
  return c.SendStatus(fiber.StatusOK)
}

func AddPhotoForm(c *fiber.Ctx) error {
  defer anc.Recover(c)
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  allSections := anc.Must(sections.GetAll()).([]sections.DataModel)
  forms.AddPhoto(allSections, nil).Render(
    context.Background(),
    c.Response().BodyWriter(),
  )
  return c.SendStatus(fiber.StatusOK)
}

