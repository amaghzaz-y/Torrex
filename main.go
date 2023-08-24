package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/server"
)

func main() {
	log.Fatalln(server.DefaultServer().Start())
	// link, err := scraper.FetchMovieLink("spider man homecoming")
	// if err != nil {
	// 	panic(err)
	// }
	// mag, err := scraper.FetchMovieMagnet(link)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(mag)
}
