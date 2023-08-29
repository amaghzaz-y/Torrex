package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET admin/room/new/:id
func (a *Api) NewRoomHanlder(c echo.Context) error {
	roomId := c.Param("id")
	if roomId == "" {
		c.Logger().Print(roomId)
		return c.String(http.StatusBadRequest, "room id is null")
	}
	room, err := a.Store.GetRoom(roomId)
	if err != nil {
		return c.String(404, "room not found")
	}
	handler, err := a.NewPipelineHandler(room)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	a.server.GET(fmt.Sprintf("/stream/%s/*", room.Id), handler)
	return c.JSON(200, fmt.Sprintf("/stream/%s/*", room.Id))
}

// GET admin/room/kill/:id
func (a *Api) KillRoomHandler(c echo.Context) error {
	roomId := c.Param("id")
	if roomId == "" {
		c.Logger().Print(roomId)
		return c.String(http.StatusBadRequest, "room id is null")
	}
	a.Torrex.StopStream(roomId)
	return c.NoContent(200)
}
