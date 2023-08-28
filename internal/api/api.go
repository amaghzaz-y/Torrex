package api

import (
	"log"
	"net/http"

	"github.com/amaghzaz-y/torrex/internal/torrex"
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

func Start() {
	api := New()
	defer api.Close()
	api.server.GET("/search/:query", api.searchHandler)
	api.server.GET("/admin/room/new/:id", api.NewRoomHanlder)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	api.server.Start(":4000")
}
