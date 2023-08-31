package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /chat/:room/new/:sender/:message
func (a *Api) newMessageHandler(c echo.Context) error {
	room := c.Param("room")
	if room == "" {
		return c.String(http.StatusBadRequest, "room id is null")
	}
	sender := c.Param("sender")
	if sender == "" {
		return c.String(http.StatusBadRequest, "sender is null")
	}
	message := c.Param("message")
	if message == "" {
		return c.String(http.StatusBadRequest, "message is null")
	}
	r := a.Torrex.Chat.ChatRoom(room)
	r.PushMessage(sender, message)
	return c.JSON(200, r.Messages())
}

// GET /chat/:room
func (a *Api) chatMessagesHandler(c echo.Context) error {
	room := c.Param("room")
	if room == "" {
		return c.String(http.StatusBadRequest, "room id is null")
	}
	r := a.Torrex.Chat.ChatRoom(room)
	return c.JSON(200, r.Messages())
}
