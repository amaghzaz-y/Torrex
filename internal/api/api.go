package api

import (
	"log"
	"net/http"

	"github.com/amaghzaz-y/torrex/internal/torrex"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Api struct {
	server *fiber.App
	*torrex.Torrex
}

func New() *Api {
	torrex := torrex.New()
	app := fiber.New()
	app.Use(cache.New())
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(helmet.New())
	app.Use(idempotency.New())
	// app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())
	return &Api{
		app,
		torrex,
	}
}

func Start() {
	api := New()
	api.server.Get("/metrics", monitor.New(monitor.Config{Title: "Torrex Metrics"}))
	api.server.Get("/search/:query", api.searchHandler)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	api.server.Listen(":4000")
}
