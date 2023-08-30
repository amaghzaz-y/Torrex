package api

import (
	"embed"
	"net/http"

	torrex "github.com/amaghzaz-y/torrex/internal/core"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Api struct {
	server *echo.Echo
	*torrex.Torrex
}

func New() *Api {
	e := echo.New()
	torrex := torrex.New()
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())
	e.Use(middleware.Decompress())
	e.Use(middleware.Logger())
	return &Api{
		e,
		torrex,
	}
}

func (api *Api) AddFS(content embed.FS) {
	api.server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(content),
		Root:       "dist/public",
		Index:      "index.html",
		HTML5:      true,
		Browse:     false,
	}))
}
func (api *Api) Start() {
	defer api.Close()
	api.server.GET("/search/:query", api.searchHandler)
	api.server.GET("/admin/room/new/:id", api.NewRoomHanlder)
	api.server.GET("/rooms", api.RoomListHandler)
	api.server.GET("/rooms/:id", api.RoomInfoHanlder)
	api.server.GET("/admin/room/kill/:id", api.KillRoomHandler)
	api.server.Start(":4000")
}
