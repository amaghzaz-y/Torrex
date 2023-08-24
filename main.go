package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/scraper"
	"github.com/amaghzaz-y/torrex/internal/torrent"
)

func main() {
	mag, err := scraper.Torrent().Magnet("Asteroid city 2023")
	if err != nil {
		panic(err)
	}
	torr := torrent.DefaultClient().NewTorrent("Asteroid city 2023", mag)
	log.Println(torr.FilePath())
	// go torr.Download()
	// for !torr.Ready() {
	// 	time.Sleep(5 * time.Second)
	// }

}
