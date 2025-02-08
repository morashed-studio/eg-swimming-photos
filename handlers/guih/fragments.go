package guih

import (
	"context"

	anc "goweb/ancillaries"
	"goweb/db/photos"
	"goweb/db/relations"
	"goweb/db/sections"
	"goweb/ui/fragments"

	"github.com/gofiber/fiber/v2"
)

func DashboardFragment(c *fiber.Ctx) error {
  c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
  sectionId := c.QueryInt("section-id", 0)

  if sectionId == 0 {
    list := anc.Must(sections.GetMain()).([]sections.DataModel)
    fragments.Dashboard(list).Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(fiber.StatusOK)
  } 

  isAlbumSection := relations.IsAlbum(sectionId)
  if isAlbumSection {
    list := anc.Must(photos.Get(sectionId)).([]photos.DataModel)
    fragments.PhotosDashboard(list).Render(context.Background(), c.Response().BodyWriter())
    return c.SendStatus(fiber.StatusOK)
  }

  ids := anc.Must(relations.GetSectionsOf(sectionId)).([]int)
  list := anc.Must(sections.Get(ids)).([]sections.DataModel)
  fragments.Dashboard(list).Render(context.Background(), c.Response().BodyWriter())
  return c.SendStatus(fiber.StatusOK)
}
