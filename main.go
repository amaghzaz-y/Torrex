package main

import (
	"github.com/amaghzaz-y/torrex/internal/server"
)

func main() {
	server.DefaultServer().Start()
	// err := torrex.TranscodeVideoToHLS("Looney/Looney.mkv", "Looney/hls/Looney.m3u8")
	// if err != nil {
	// 	panic(err)
	// }
	// c := torrex.NewClient()
	// if err := c.RegisterTorrentFile("assets/looney.torrent"); err != nil {
	// 	log.Fatalln(err)
	// }
	// c.DownloadFiles()
	// for {
	// 	if c.IsDownloadComplete() {
	// 		break
	// 	}
	// 	c.LogInfo()
	// 	time.Sleep(5 * time.Second)
	// }
	// log.Println("finished download")
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
