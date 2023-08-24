package main

import (
	"log"
	_ "net/http/pprof"
	"time"

	"github.com/amaghzaz-y/torrex/internal/scraper"
	"github.com/amaghzaz-y/torrex/internal/server"
	"github.com/amaghzaz-y/torrex/internal/streamer"
	"github.com/amaghzaz-y/torrex/internal/torrent"
)

func main() {
	mag, err := scraper.Torrent().Magnet("Asteroid city 2023")
	if err != nil {
		panic(err)
	}
	torr := torrent.DefaultClient().NewTorrent("Asteroid city 2023", mag)
	go torr.Download()
	for !torr.Ready() {
		time.Sleep(5 * time.Second)
	}
	handler := streamer.NewStreamer().BootstrapStream("Asteroid city 2023", torr.FilePath(), torr.UdpPort())
	server := server.DefaultServer()
	server.Handle("/*", handler)
	log.Fatalln(server.Start())
}
