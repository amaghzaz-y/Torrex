package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/server"
)

func main() {
	log.Println(server.DefaultServer().Start())
}
