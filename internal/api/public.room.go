package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /rooms
func (a *Api) roomListHandler(c echo.Context) error {
	return c.JSON(200, a.Rooms())
}

// GET /rooms/:id

func (a *Api) roomInfoHanlder(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "room id is null")
	}
	room, err := a.Store.GetRoom(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, room)
}
