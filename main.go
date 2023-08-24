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
	log.Println("torrent found")

	go torr.Download()
	log.Println("torrent download started")
	for !torr.Ready() {
		log.Println("waiting for stream to be ready")
		time.Sleep(5 * time.Second)
	}
	log.Println("stream is ready")
	handler := streamer.NewStreamer().BootstrapStream("Asteroid city 2023", torr.FilePath(), torr.UdpPort())
	log.Println("stream is boostrapped")
	server := server.DefaultServer()
	server.Handle("/*", handler)
	log.Println("server started")
	log.Fatalln(server.Start())
}
