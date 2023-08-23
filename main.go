package main

import (
	"log"

	"github.com/amaghzaz-y/torrex/pkg/scraper"
)

func main() {
	// c := torrex.NewClient()
	// if err := c.RegisterTorrentFile("looney.torrent"); err != nil {
	// 	log.Fatalln(err)
	// }
	// c.DownloadFiles()
	// for {
	// 	if c.IsDownloadComplete() {
	// 		break
	// 	}
	// }
	// log.Println("finished download")
	link, err := scraper.FetchMovieLink("spider man")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(link)
}
