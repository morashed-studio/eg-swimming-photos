package section

import (
	anc "goweb/ancillaries"
	"goweb/db/relations"
	"goweb/db/sections"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddSection(c *fiber.Ctx) error {
  defer anc.Recover(c)
  body := new(AddSectionBody)
  if err := c.BodyParser(body); err != nil {
    return c.Status(fiber.StatusBadRequest).SendString(err.Error())
  }

  newSection := sections.DataModel { Title: body.Title }
  anc.Must(nil, sections.Add([]sections.DataModel{ newSection }))

  if body.Parent == "none" {
    return c.SendStatus(fiber.StatusOK)
  }

  newRelation := relations.DataModel { 
    Parent: anc.Must(strconv.Atoi(body.Parent)).(int),
    Child: anc.Must(sections.GetId(body.Title)).(int),
  }
  anc.Must(nil, relations.Add([]relations.DataModel{ newRelation }))

  return c.SendStatus(fiber.StatusOK)
}

