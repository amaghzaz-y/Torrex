package main

import (
	"log"
	"time"

	torrex "github.com/amaghzaz-y/torrex/internal"
)

func main() {
	c := torrex.NewClient()
	if err := c.RegisterTorrentFile("assets/looney.torrent"); err != nil {
		log.Fatalln(err)
	}
	c.DownloadFiles()
	for {
		if c.IsDownloadComplete() {
			break
		}
		c.LogInfo()
		time.Sleep(5 * time.Second)
	}
	log.Println("finished download")
	// link, err := scraper.FetchMovieLink("spider man")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// link, err = scraper.FetchMovieMagnet(link)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(link)
}
