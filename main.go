package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/api"
)

func main() {
	// mag, err := scraper.Torrent().Magnet("Asteroid city 2023")
	// if err != nil {
	// 	panic(err)
	// }
	// torr := torrent.DefaultClient().NewTorrent("Asteroid city 2023", mag)
	// log.Println("torrent found")

	// go torr.Download()
	// log.Println("torrent download started")
	// for !torr.Ready() {
	// 	log.Println("waiting for stream to be ready :", torr.Completion())
	// 	time.Sleep(5 * time.Second)
	// }
	// log.Println("stream is ready")
	// handler := streamer.NewStreamer().BootstrapStream("Asteroid city 2023", torr.FilePath(), torr.UdpPort())
	// log.Println("stream is boostrapped")
	log.Println("server started")
	api.Start()
}
