package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GET admin/room/new/:id
func (a *Api) NewRoomHanlder(c *fiber.Ctx) error {
	roomId := c.Params("id")
	if roomId == "" {
		return c.Status(http.StatusBadRequest).SendString("room id is null")
	}
	room, err := a.Store.GetRoom(roomId)
	if err != nil {
		return c.Status(404).SendString("room not found")
	}
	handler := a.NewPipelineHandler(room)
	a.server.Get(fmt.Sprintf("/%s", room.Id), handler)
	return c.SendString(fmt.Sprintf("/%s", room.Id))
}
