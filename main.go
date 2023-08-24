package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/amaghzaz-y/torrex/internal/scraper"
)

func main() {
	// log.Fatalln(server.DefaultServer().Start())
	link, err := scraper.Info().FetchMovieInfoLink("red notice")
	if err != nil {
		panic(err)
	}
	fmt.Println(link)
	info, err := scraper.Info().FetchMovieInfo(link)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", info)
}
