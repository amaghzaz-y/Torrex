package main

import (
	"embed"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/api"
)

//go:embed all:dist/public
var content embed.FS

func main() {
	api := api.New()
	api.AddFS(content)
	api.Start()
}
